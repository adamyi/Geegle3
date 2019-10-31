import re
import os
import sys

from gunicorn.app.wsgiapp import run
from whoosh.index import create_in, open_dir

os.system("python -m compileall /app/chals/web/search/image.binary.runfiles/")
print("Initiating")
# ix = create_in("/index", WebSchema)
try:
    ix = open_dir("index")
except:
    os.system(
        "wget --progress=dot:giga https://geegle-index.s3-ap-southeast-2.amazonaws.com/index.tar && tar xvf index.tar && \\rm -f index.tar"
    )
print("Initiaing complete")

sys.argv[0] = re.sub(r'(-script\.pyw?|\.exe)?$', '', sys.argv[0])
sys.argv.append("--workers=8")
sys.argv.append("app:app")
sys.argv.append("-b")
sys.argv.append("0.0.0.0:80")
sys.exit(run())
