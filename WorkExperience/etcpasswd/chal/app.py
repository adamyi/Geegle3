#!/usr/bin/env python

import os
import time
import sys

for c in "Loading...\n":
    time.sleep(0.2)
    print(c, end='')
    sys.stdout.flush()

print("Welcome to sysadmin control system")
username = input("Username: ")
password = input("Password: ")

if username == "root" and password == "cocacola":
    print("Welcome... enjoy your shell")
    print("$ ")
    sys.stdout.flush()
    os.system("sh")
