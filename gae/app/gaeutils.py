# GAE Utils
import json

def json_encode(value):
  return json.dumps(value)

def json_decode(text):
  return json.loads(text)
