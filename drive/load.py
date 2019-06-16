import test_pb2

KEY_SIZE = 64

# f = open('a', 'rb')
f = open('b', 'rb')
b = test_pb2.FileRequest()
c = f.read()[KEY_SIZE:]
# print(c)
b.ParseFromString(c)
print(b)
