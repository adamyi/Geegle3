import re
import sys

from gunicorn.app.wsgiapp import run

sys.argv[0] = re.sub(r'(-script\.pyw?|\.exe)?$', '', sys.argv[0])
sys.argv.append("app:app")
sys.argv.append("-b")
sys.argv.append("0.0.0.0:80")
sys.exit(run())
