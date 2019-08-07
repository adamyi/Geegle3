#!/usr/bin/env python3

import time
import random
import signal



nums = []
for i in range(1000):
    nums.append(random.randint(1,1000000))

print("I need some things sorted ASAP!")
print("make sure you give them all back on their own line, exactly how I gave them to you:")
time.sleep(5)


for elem in nums:
    print(elem)
print("What is your response (you have 10 seconds):")

nums.sort()

def handler(a, b):
    print('You took too long')
    exit(0)

signal.signal(signal.SIGALRM, handler)
signal.alarm(10)


responses = []
for i in range(len(nums)):
    try:
        responses.append(int(input()))
    except ValueError:
        print("Make sure you input numbers!!")

signal.alarm(0) # Disable alarm

if nums == responses:
    try:
        fh = open("/flag", "r")
    except FileNotFoundError:
        print("FLAG{DEBUG_FLAG}")
        exit(1)
    print(fh.read())
    fh.close()
    exit(0)

print("make sure you get it right next time!")
