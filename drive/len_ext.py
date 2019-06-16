import httplib, urlparse, sys, urllib
from pymd5 import md5, padding
from pwn import *

curhash = "fc01d71f668179d67c0217dc3a48d179"
f = open("./a")
msg = f.read()
f.close()

mlen = len(msg)
print("Original length in bytes: %d" % mlen)
bits = (mlen + len(padding(mlen*8)))*8
print("Original length in bits (with padding): %d" % bits)

h = md5(state=curhash.decode("hex"), count=bits)
x = "\x80" + "\xb8\x17\x00\x15\x00\x00\x00\x00" + "\x1a\x06pwned1" + "\x0a\x06pwned2"
h.update(x)

#generate new hash and url
newhash = h.hexdigest()
padding = padding(mlen*8)
print("Padding found at end of original instructions: "+padding)
msg = msg + padding + x

print("New hash to be inserted: "+newhash)

f = open("./b", "w")
f.write(msg)
f.close()

