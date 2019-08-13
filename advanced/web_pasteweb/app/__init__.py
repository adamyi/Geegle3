import re
import datetime
import sqlite3
import uuid
import json
import base64
from flask import (
  Flask,
  render_template_string,
  request,
  render_template,
  current_app,
  flash,
  redirect,
  url_for,
  session,
  make_response,
  Response,
)
from talisman import Talisman
from lxml import etree

app = Flask(__name__)
app.secret_key = "haha"

csp = {
    'object-src': '\'none\'',
    'base-uri': '\'none\'',
    'script-src': [
      '\'unsafe-eval\'',
      '\'strict-dynamic\'',
      '\'sha384-tsQFqpEReu7ZLhBV2VZlAu7zcOV+rXbYlF2cqB8txI/8aZajjp4Bqd+V6D5IgvKT\'', # jquery
      '\'sha384-ELH09WGRUcBpRT6iHTekFB2YBCT9kFMsKG4Y9LUAevHjihu8Otri8Sm01QgXOTht\'', # dompurify
      '\'sha384-cC2m6yBYbHGgI8+RaJwvRVIG/+yU8d8oMCriPmy9bUONAcpdtJHSbR7xHq6Z9m7P\'', # feedback.js
    ],
}

talisman = Talisman(
  app,
  content_security_policy=csp,
  content_security_policy_nonce_in=['script-src'],
  content_security_policy_report_uri="/api/bugreport/csp",
  legacy_content_security_policy_header=False,
  session_cookie_secure=False,
  force_https=False,
  strict_transport_security=False
)

LINK_REGEX = re.compile(r'(https?://.*?)["\' `]')

def xmlcheck(xml, level):
  if level >= 3:
    raise ValueError("Too many levels for DTD!")
  badwords = ["file://", "etc", "passwd"]
  if any(word in xml for word in badwords):
    raise ValueError("Bad word " + word + " detected.")

  for link in LINK_REGEX.findall(xml):
    r = ""
    try:
      r = requests.get(link)
    except:
      raise ValueError("Failed to request data from " + link)
    xmlcheck(r.test, level + 1)

@app.route("/")
def home():
  return render_template("home.html")

@app.route("/new", methods=["GET", "POST"])
def newpaste():
  if request.method == "POST":
    title = request.form["title"]
    contents = request.form["contents"]
    post_id = str(uuid.uuid4())
    try:
      with sqlite3.connect("pasteweb.db") as conn:
        c = conn.cursor()
        c.execute("INSERT INTO pastes (id, title, contents) VALUES (?, ?, ?)", (post_id, title, contents))
        conn.commit()
    except Exception as e:
      print(e)
      flash("Something went wrong while creating your paste, try again later.", "danger")
      return render_template("new_paste.html")
    else:
      flash("Created paste successfully!", "success")

      return redirect("/paste/"+post_id)
  return render_template("new_paste.html")


@app.route("/paste/<paste_id>")
def viewpaste(paste_id):
  try:
    with sqlite3.connect("pasteweb.db") as conn:
      c = conn.cursor()
      c.execute('SELECT title, contents FROM pastes WHERE id=?', (paste_id,))
      paste = c.fetchone()
  except Exception as e:
    flash("Paste Not Found", "danger")
    return render_template("home.html"), 404

  rsp = make_response(render_template("paste.html", title=base64.b64encode(paste[0]), content=base64.b64encode(paste[1])))
  rsp.set_cookie('pasteweb_debug', 'viewable by pasteweb developer only')
  return rsp

@app.route("/api/bugreport/<bugtype>", methods=["POST"])
def bugreport(bugtype):
  if bugtype != "csp" and bugtype != "js":
    return "Wrong bugtype!"
  if bugtype == "js":
    if request.content_type == "application/json":
      return make_response(("Bug Received", {"Accept": "application/json"}))
    return make_response(("Content-Type unknown", 400, {"Accept": "application/json"}))
  elif bugtype == "csp":
    if request.content_type == "application/csp-report":
      return make_response(("Bug Received", {"Accept": "application/csp-report, application/xml"}))
    if request.content_type == "application/xml":
      xml = request.data
      try:
        xmlcheck(xml, 0)
      except Exception as e:
        return make_response((e.message, 500, {"Accept": "application/csp-report, application/xml"}))

      try:
        #parser = etree.XMLParser(dtd_validation=True, load_dtd=True, no_network=False, huge_tree=True)
        parser = etree.XMLParser(encoding="utf-8", no_network=False, huge_tree=True)
        etree.fromstring(xml.encode("utf-8"), parser=parser)
      except Exception as e:
        return make_response((e.message, 500, {"Accept": "application/csp-report, application/xml"}))
      return make_response(("Bug Received", {"Accept": "application/csp-report, application/xml"}))
    return make_response(("Content-Type unknown", 400, {"Accept": "application/csp-report, application/xml"}))



if __name__ == '__main__':
  app.run(debug=True, host='0.0.0.0', port=80)
