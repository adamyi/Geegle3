#!/usr/bin/env python3

import time
import random

MAX_TIME = 1000 * 1

nums = []
for i in range(1000):
    nums.append(random.randint(1,1000000))

print("I need some things sorted ASAP!")
print("make sure you give them all back on their own line, exactly how I gave them to you:")
time.sleep(5)


for elem in nums:
    print(elem)
print("What is your response:")

startTime = time.time()*1000.0
nums.sort()

responses = []
for i in range(len(nums)):
    try:
        responses.append(int(input()))
    except ValueError:
        print("Make sure you input numbers!!")
endTime = time.time()*1000.0

if endTime - startTime > MAX_TIME:
    print("No, I need them faster than that!\nThis is useless now!!")
    exit(0)

if nums == responses:
    try:
        fh = open("/flag", "r")
    except FileNotFoundError:
        print("FLAG{DEBUG_FLAG}")
        exit(1)
    print(f.read())
    fh.close()
    exit(0)

print("make sure you get it right next time!")
