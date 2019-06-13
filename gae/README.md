# gae (Geegle App Engine)

Like Amazon Lambda.

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
