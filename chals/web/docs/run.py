import re
import sys
import os

from gunicorn.app.wsgiapp import run

os.system("python -m compileall /app/chals/web/docs/image.binary.runfiles/")
sys.argv[0] = re.sub(r'(-script\.pyw?|\.exe)?$', '', sys.argv[0])
sys.argv.append("--workers=4")
sys.argv.append("app:app")
sys.argv.append("-b")
sys.argv.append("0.0.0.0:80")
sys.exit(run())
