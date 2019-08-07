from flask import Flask

HOST = 'apps.geegle.org'
SERVER_NAME = 'Geegle Frontend'

class gaeFlask(Flask):
  def process_response(self, response):
    super(gaeFlask, self).process_response(response)
    response.headers['Server'] = SERVER_NAME
    return(response)

app = gaeFlask(__name__)
app.config['SERVER_NAME'] = HOST

from app import routes
print("gae init")
