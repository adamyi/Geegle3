from flask import Flask, Response, request, make_response, render_template
import requests
import yaml
import json
import base64
from urlparse import urlparse
from asteval import Interpreter
from mss import MSSClient
import gaeutils

HOST = 'apps.geegle.org:8056'

#TODO: increase CACHE_EXPIRES to 5min/10min
CACHE_EXPIRES = 60 # 1 min

app = Flask(__name__)
app.config['SERVER_NAME']=HOST

# TODO: don't hard code this
# use auth service instead
mssc = MSSClient("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFwcGVuZ2luZUBzZXJ2aWNlcy5nZWVnbGUub3JnIn0.L3B8lu7bMUx6YUNDZn1ltYYrNwsgxjcYmFUOH9ZFYw0")

def strip_scheme(url):
  parsed = urlparse(url)
  scheme = "%s://" % parsed.scheme
  return parsed.geturl().replace(scheme, '', 1)

def error(code, message, debug=""):
  return render_template("error.html", code=code, message=message, debug=debug), code

@app.route('/',  subdomain = "manage",  methods=['GET', 'POST'])
def manage():
  return "this is manage portal"

@app.route('/edit',  subdomain = "manage",  methods=['GET', 'POST'])
def edit():
  # TODO: permission check
  userapp = request.args.get('app')
  if not userapp:
    return error(400, "wrong parameter")
  if request.method == "GET":
    return render_template("index.html", code=mssc.get("code/" + userapp))
  if request.method == "POST":
    code = request.form["code"]
    decoded = yaml.safe_load(code)
    if "urls" not in decoded:
      return error(400, "missing urls")
    if "default_handler" not in decoded:
      return error(400, "missing default_handler")
    for url in decoded['urls']:
      if not url.startswith("/"):
        return error(400, "%s must start with /" % url)
    for url, code in decoded['urls'].items():
      mssc.set("cache/" + userapp + '.' + HOST + url, code, CACHE_EXPIRES)
    mssc.set("code/" + userapp, request.form["code"])
    return "Updated"
  

@app.route('/', subdomain="<userapp>", defaults={'path': ''}, methods=['GET', 'POST'])
@app.route('/<path:path>', subdomain = "<userapp>", methods=['GET', 'POST'])
def runapp(path, userapp):
  code = mssc.get("cache/" + strip_scheme(request.base_url))
  if code is None or code == "":
    sitecode = mssc.get("code/" + userapp)
    if sitecode is None or sitecode == "":
      return error(404, "The requested URL %s was not found on this server." % request.path)
    decoded = yaml.safe_load(mssc.get("code/" + userapp))
    if request.path in decoded['urls']:
      code = decoded['urls'][request.path]
    else:
      code = decoded['default_handler']
    mssc.set("cache/" + request.base_url, code, CACHE_EXPIRES)

  aeval = Interpreter(usersyms={'request': request, 'gaeutils': gaeutils})
  aeval(code)
  if len(aeval.error) > 0:
    dbg = base64.b64encode(json.dumps([err.get_error() for err in aeval.error]))
    return error(500, "That was an error. Please try again later.", debug=dbg)

  return aeval.symtable['ret_body'], aeval.symtable['ret_status']

if __name__ == '__main__':
  app.run(debug=False, host='0.0.0.0', port=80)
