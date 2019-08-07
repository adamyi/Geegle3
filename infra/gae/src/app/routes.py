from flask import Response, request, make_response, render_template
import requests
import yaml
from urlparse import urlparse
from mss import MSSClient
from app import app
from gaerunner import runCode, flattenRequest
import gaeutils

CACHE_EXPIRES = 300 # 5 min
REQUEST_TIMEOUT = 3 # seconds

# TODO: don't hard code this
# use auth service instead
mssc = MSSClient("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFwcGVuZ2luZUBzZXJ2aWNlcy5nZWVnbGUub3JnIn0.L3B8lu7bMUx6YUNDZn1ltYYrNwsgxjcYmFUOH9ZFYw0")

def strip_scheme(url):
  parsed = urlparse(url)
  scheme = "%s://" % parsed.scheme
  return parsed.geturl().replace(scheme, '', 1)

def error(code, message="That was an error. Please try again later.", debug=""):
  return render_template("error.html", code=code, message=message, debug=debug), code

@app.route('/',  subdomain = "manage",  methods=['GET', 'POST'])
def manage():
  return "this is manage portal"

@app.route('/edit',  subdomain = "manage",  methods=['GET', 'POST'])
def edit():
  # TODO: permission check
  userapp = request.args.get('app')
  if not userapp:
    return error(400, message="wrong parameter")
  if request.method == "GET":
    return render_template("index.html", code=mssc.get("code/" + userapp))
  if request.method == "POST":
    code = request.form["code"]
    decoded = yaml.safe_load(code)
    if "urls" not in decoded:
      return error(400, message="missing urls")
    if "default_handler" not in decoded:
      return error(400, message="missing default_handler")
    for url in decoded['urls']:
      if not url.startswith("/"):
        return error(400, message="%s must start with /" % url)
    for url, code in decoded['urls'].items():
      mssc.set("cache/" + userapp + '.' + app.config['SERVER_NAME'] + url, code, CACHE_EXPIRES)
    mssc.set("code/" + userapp, request.form["code"])
    return "Updated"
  

@app.route('/', subdomain="<userapp>", defaults={'path': ''}, methods=['GET', 'POST'])
@app.route('/<path:path>', subdomain = "<userapp>", methods=['GET', 'POST'])
def runapp(path, userapp):
  try:
    code = mssc.get("cache/" + strip_scheme(request.base_url))
    if code is None or code == "":
      sitecode = mssc.get("code/" + userapp)
      if sitecode is None or sitecode == "":
        return error(404, message="The requested URL %s was not found on this server." % request.path)
      decoded = yaml.safe_load(mssc.get("code/" + userapp))
      if request.path in decoded['urls']:
        code = decoded['urls'][request.path]
      else:
        code = decoded['default_handler']
      mssc.set("cache/" + request.base_url, code, CACHE_EXPIRES)

    result = runCode({'request': flattenRequest(request), 'gaeutils': gaeutils}, code, REQUEST_TIMEOUT)
    if not result:
      return error(502, message="The server didn't return anything.")
    if type(result) == unicode:
      return error(500, debug=result.encode('utf-8'))
    if result['errpage']:
      rsp = make_response(error(result['status'], message=result['body']))
    else:
      rsp = make_response(result['body'], result['status'])
    for k, v in result['headers'].items():
      rsp.headers[k] = v
    return rsp
  except Exception as e:
    return error(500, debug=str(e))
