# sffe (static-on-sstable / static file front-end)

Just like gstatic.com

**Feel free to use this infra in other services**

## How-To
```
POST http://127.0.0.1:8067/api/store/
{"filename":"test.txt", "content":"aGVsbG8gd29ybGQhCg==", "flags":[{"name":"pa","value":"1"},{"name":"lol","value":"haha"}]}
```

```
GET http://127.0.0.1:8067/s/098f6bcd4621d373cade4e832627b4f6/lol=haha/pa=1/test.txt

098f6bcd4621d373cade4e832627b4f6 is md5 of service name test
Flags KV are alphabetically sorted (URL is also returned by storage api)
```
