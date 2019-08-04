#!/usr/bin/env python2

from pwn import *

p = remote("localhost", 8999)#process('./gen.py')

p.recvline()
p.recvline()

nums = []
while True:
    try:
        line = p.recvline()
        num = int(line)
        nums.append(num)
    except:
        break

nums.sort()
for i in nums:
    p.sendline(str(i))

p.interactive()
