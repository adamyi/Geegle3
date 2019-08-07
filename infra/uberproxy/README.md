# ÜberProxy/BeyondCorp
Reverse Proxy for \*.corp.geegle.org + centralized authentication

We have Zero-Trust network lmao

## vulnerability

It directly forwards OPTIONS request without authenticating for CORS preflight requests.

If one internal service treats OPTIONS request as normal GET request (e.g. by default,
Golang doesn't care what request method it is), this might lead to unauthenticated access.

This itself is not a vulnerability as long as all internal services react to OPTIONS
correctly.

TODO(adamyi): add a challenge to use this vulnerability

## author
adamyi
