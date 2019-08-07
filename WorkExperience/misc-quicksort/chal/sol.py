#!/usr/bin/env python2

from pwn import *

#p = remote("localhost", 8999)#process('./gen.py')
p = process("python cli-relay.py eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY2RAZ2VlZ2xlLm9yZyIsInNlcnZpY2UiOiJ1YmVycHJveHlAc2VydmljZXMuZ2VlZ2xlLm9yZyIsImV4cCI6MTU2NzczNzExM30.B7-MHoEPrU-GvEQ0tx1dHtfcpBUlcZOH_3F3z1PuNvA quicksort.corp.geegle.org", shell=True)

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
