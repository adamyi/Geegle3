import os
import json
import time
import urllib
import requests

from flask import render_template, Flask, request, session, send_file

from whoosh.analysis import NgramTokenizer
from whoosh.fields import SchemaClass, TEXT, KEYWORD, ID
from whoosh.index import create_in, open_dir
from whoosh.qparser.default import MultifieldParser
from whoosh.qparser import OrGroup
from whoosh.query import FuzzyTerm
from whoosh.highlight import SCORE, HtmlFormatter, ContextFragmenter


class WebSchema(SchemaClass):
    title = TEXT(stored=True, analyzer=NgramTokenizer(minsize=2, maxsize=10))
    url = TEXT(stored=True)
    content = TEXT(stored=True, analyzer=NgramTokenizer(minsize=2, maxsize=10))


class MyFuzzyTerm(FuzzyTerm):
    def __init__(self,
                 fieldname,
                 text,
                 boost=1.0,
                 maxdist=1,
                 prefixlength=1,
                 constantscore=True):
        super(MyFuzzyTerm, self).__init__(fieldname, text, boost, maxdist,
                                          prefixlength, constantscore)


import flag

app = Flask(__name__)

app.config.update({
    "SECRET_KEY": "supersecretkey",
})

NO_PROXY = {
    'no': 'pass',
}

#indexpath = "index.%d" % os.getpid()
#os.system("cp -r index " + indexpath)
indexpath = "index"
ix = open_dir(indexpath)


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/flag", methods=["GET", "POST", "FLAG"])
def get_flag():
    #if request.method != "FLAG":
    #    return "You sent a %s request. Please send a FLAG request to view flag!" % request.method
    ip = request.remote_addr
    if ip == "127.0.0.1":
        return flag.FLAG
    else:
        return "only 127.0.0.1 can view flag"


@app.route("/robots.txt")
def robots():
    return render_template("robots.html"), 200, {'Content-Type': 'text/plain'}


@app.route("/admin")
def admin():
    return render_template("admin.html")


# @app.route("/admin/geegle3")
# def source():
#   return send_file(__file__)


@app.route("/admin/crawl", methods=["POST"])
def crawl():
    form = request.form
    if not form['url'].startswith('http'):
        return "Invalid URL"
    if form["url"]:
        print(form['url'])
        session = requests.Session()
        session.trust_env = False
        t = str(session.get(form['url'], proxies=NO_PROXY).text)
        return "sorry you don't have permission to add to the index but here's your page:\n" + t, 200, {
            'Content-Type': 'text/plain'
        }
        #asyncio.ensure_future(check_url(form["url"]))
    #return "It's been added to queue and should be indexed in a minute if things check out."


@app.route("/search", methods=["GET"])
def search():
    start_time = time.time()
    form = request.form
    qstr = request.args.get('q')
    page = int(request.args.get('p', "1"))
    parser = MultifieldParser(['title', 'content'],
                              schema=ix.schema,
                              group=OrGroup)
    #termclass=MyFuzzyTerm)
    query = parser.parse(qstr)
    notes = []
    with ix.searcher() as searcher:
        corrected = searcher.correct_query(query, qstr)
        results = searcher.search(query)
        rel = results.estimated_length()
        if corrected.string.lower() != qstr.lower():
            crs = searcher.search(corrected.query)
            if crs.estimated_length() > rel:
                notes.append("Did you mean: " + corrected.string)
        results = searcher.search_page(query, page, terms=True)
        my_cf = ContextFragmenter(maxchars=20, surround=30, charlimit=256)
        results.order = SCORE
        results.fragmenter = my_cf
        # results.formatter = HtmlFormatter()
        rsp = [{
            "url": item["url"],
            "content": item.highlights("content"),
            "title": item["title"]
        } for item in results]
    # return json.dumps(rsp)
    # print(json.dumps(rsp))
    if rel == 0:
        notes.append("Sorry, no result for your query")
    else:
        elapsed_time = time.time() - start_time
        notes.append("%d results found in %.2f seconds" % (rel, elapsed_time))
    return render_template(
        "result.html",
        result=rsp,
        query=qstr,
        notes=notes,
        nextpage=page + 1,
        urlquery=urllib.quote_plus(qstr))


def check_url(url):
    #task = asyncio.ensure_future(start_browser(url))
    #asyncio.sleep(10)
    #task.cancel()
    #try:
    #    await task
    #except asyncio.CancelledError:
    pass


#def start_browser(url):
#    print("browse " + url)
#    browser = await launch(**chromium_path(), args=["--no-sandbox"])
#    result = {}
#    try:
#        page = await browser.newPage()
#        await page.goto(url, timeout=5000)
#        title = await page.title()
#        content = await page.evaluate('() => document.body.innerText')
#        add_to_index(url, title, content)
#    finally:
#        await browser.close()


def add_to_index(url, title, content):
    print("index " + url)
    writer = ix.writer()
    writer.add_document(title=title, url=url, content=content)
    writer.commit()


def chromium_path():
    # return {"executablePath": "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"}
    if os.path.isfile("/usr/bin/chromium"):
        return {"executablePath": "/usr/bin/chromium"}
    return {}


if __name__ == "__main__":
    app.run(port=5055, debug=True)
