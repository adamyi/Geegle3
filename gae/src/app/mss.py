# mss library

import requests
# import httplib as http_client
# http_client.HTTPConnection.debuglevel = 1

class MSSClient():
  # TODO: don't hard code jwt
  # use our auth service instead
  def __init__(self, jwt, host="http://mss"):
    self.jwt = jwt
    self.host = host

  def get(self, key):
    return requests.get(self.host + "/" + key, headers={"X-Geegle-JWT":self.jwt}).text

  def set(self, key, value, expires=None):
    headers = {"X-Geegle-JWT":self.jwt}
    if expires:
      headers["X-MSS-Expires"] = str(expires)
    requests.post(self.host + "/" + key, headers=headers, data=value)
