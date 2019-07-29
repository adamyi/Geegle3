#!/usr/bin/env python3

import os
import string
import random


MB = 300
# MBs to write, in bytes
MB_SIZE = 1024 * 1024 * MB

def updateProgress(i):
    global MB
    print("\rProgress: [", end='')
    total = int(MB/5)
    progress = int(i / 1024 / 1024 / 5)
    for j in range(progress):
        print("=", end='')
    for j in range(total - progress):
        print(" ", end='')
    print("]", end='\r')


fh = open("output.log", "w+")

# Write a garbage file with the flag in the middle
i = 0
updateProgress(i)
while i < MB_SIZE * 0.5:
    fh.write(random.choice(string.ascii_letters))
    i += 1
    if i % (1024 * 1024) == 0:
        updateProgress(i)

try:
    flagFile = open("/flag", "r")
    fh.write(flagFile.read())
    flagFile.close()
except:
    fh.write("FLAG{DEBUG_FLAG}")

while i < MB_SIZE:
    fh.write(random.choice(string.ascii_letters))
    i += 1
    if i % (1024 * 1024) == 0:
        updateProgress(i)

fh.close()
