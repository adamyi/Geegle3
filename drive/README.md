# drive (Geegle Drive)

PoC only for now.

## Solution
You need to understand how Protobuf works
(how to encode by hand)

Then it's just Hash Length Extension Attack

## 48000
It took me hours to make this work and the number also looks really nice.

## PoC
```
$ protoc test.proto --python_out=.
$ protoc test.proto --go_out=.
$ python test.py

\x17/user=test/s=1/test.txt\x15\x00\x00\x00\x1aK[drive@services.geegle.org] test is accessing this file (X-Geegle-Debug: ).

$ python len_ext.py
Original length in bytes: 44040951
Original length in bits (with padding): 352327680
Padding found at end of original instructions: \x80\xb8\x17\x00\x15\x00\x00\x00\x00
New hash to be inserted: 7caec8fb5aa7a8f9a2368e5d59f47853

$ go run load.go test.pb.go
filename:"pwned2" log_info:"pwned1"

$ md5 b
MD5 (b) = 7caec8fb5aa7a8f9a2368e5d59f47853

 ```
