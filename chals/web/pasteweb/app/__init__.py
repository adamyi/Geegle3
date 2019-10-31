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
from flask_talisman import Talisman
from lxml import etree
import jwt

app = Flask(__name__)
app.secret_key = "haha"

csp = {
    'object-src':
    '\'none\'',
    'base-uri':
    '\'none\'',
    'script-src': [
        '\'unsafe-eval\'',
        '\'strict-dynamic\'',
        # NOTE(adamyi@): idk why but firefox doesn't like this... switch to nonce for them as well... (chrome works fine)
        #'\'sha384-vk5WoKIaW/vJyUAd9n/wmopsmNhiy+L2Z+SBxGYnUkunIxVxAv/UtMOhba/xskxh\'', # jquery
        #'\'sha384-ELH09WGRUcBpRT6iHTekFB2YBCT9kFMsKG4Y9LUAevHjihu8Otri8Sm01QgXOTht\'', # dompurify
    ],
}

talisman = Talisman(
    app,
    content_security_policy=csp,
    content_security_policy_nonce_in=['script-src'],
    content_security_policy_report_uri=
    "https://bugreport.corp.geegle.org/api/bugreport/csp",
    legacy_content_security_policy_header=False,
    session_cookie_secure=False,
    force_https=False,
    strict_transport_security=False)



@app.route("/")
def home():
    return render_template("home.html")


@app.route("/new", methods=["GET", "POST"])
def newpaste():
    if request.method == "POST":
        title = request.form["title"]
        content = request.form["content"]
        if len(title) > 35:
            flash("Title cannot be longer than 35 chars!", "danger")
            return render_template("new_paste.html")
        post_id = str(uuid.uuid4())
        try:
            with sqlite3.connect("/pasteweb.db") as conn:
                c = conn.cursor()
                c.execute(
                    "INSERT INTO pastes (id, title, contents) VALUES (?, ?, ?)",
                    (post_id, title, content))
                conn.commit()
        except Exception as e:
            print(e)
            flash(
                "Something went wrong while creating your paste, try again later.",
                "danger")
            return render_template("new_paste.html")
        else:
            flash("Created paste successfully!", "success")

            return redirect("/paste/" + post_id)
    return render_template("new_paste.html")


@app.route("/paste/<paste_id>")
def viewpaste(paste_id):
    try:
        with sqlite3.connect("/pasteweb.db") as conn:
            c = conn.cursor()
            c.execute('SELECT title, contents FROM pastes WHERE id=?',
                      (paste_id, ))
            paste = c.fetchone()
    except Exception as e:
        flash("Paste Not Found", "danger")
        return render_template("home.html"), 404

    if paste is None:
        flash("Paste Not Found", "danger")
        return render_template("home.html"), 404

    rsp = make_response(
        render_template(
            "paste.html", title=(paste[0]),
            content=base64.b64encode(paste[1])))
    print(
        jwt.decode(request.headers.get('X-Geegle-JWT'),
                   'superSecretJWTKEY')['username'])
    if jwt.decode(request.headers.get('X-Geegle-JWT'), 'superSecretJWTKEY'
                  )['username'] == 'xssbot+pasteweb@services.geegle.org':
        rsp.set_cookie('pasteweb_debug', 'GEEGLE{JAO34OADS81HI}')
    else:
        rsp.set_cookie('pasteweb_debug', 'viewable by pasteweb developer only')
    return rsp


if __name__ == '__main__':
    app.run(debug=True, host='0.0.0.0', port=80)
