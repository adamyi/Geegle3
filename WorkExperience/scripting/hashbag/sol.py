#!/usr/bin/env python3

import os
import hashlib

pwfile = open('password.md5', 'r')
hash = pwfile.read().rstrip()
print("Hash: \t\t[" + hash + "]")
pwfile.close()

fh = open('wordlist.txt', 'r')

i = 0
for line in fh:
    line = line.rstrip()
    if hash == hashlib.md5(line.encode()).hexdigest():
        print("Password: \t[", end='')
        print(line, end='')
        print("]")
        break
    i += 1
fh.close()

