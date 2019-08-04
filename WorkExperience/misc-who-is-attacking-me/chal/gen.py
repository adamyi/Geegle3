#!/usr/bin/env python3
import time
import random
import signal
import json
from collections import defaultdict

target = "192.152.10.52"
IPS = ["172.41.15.132",
    "50.37.251.207",
    "17.214.11.56",
    "233.172.200.44",
    "11.165.2.220",
    "137.46.136.134",
    "205.103.91.30",
    "229.27.227.65",
    "195.80.95.56",
    "87.114.56.91",
    "127.79.181.98",
    target]
SIZES = [2**i for i in range(0, 10)]

print("Hello intern. We are under attack!")
print("We need your help to identify who is attacking us...")
print(f"Who is sending the most data to {target}")

time.sleep(2)
print("Ready?")
time.sleep(1)

logs = []

MAX_target = "empty"
MAX_target_data = 0
target_data = defaultdict(int)

for i in range(0, 500):
    data = {"SRC": random.choice(IPS), "DST": random.choice(IPS), "data_size": random.choice(SIZES)}

    if data["DST"] == target:
        target_data[data["SRC"]] = target_data[data["SRC"]] + int(data["data_size"])
        if target_data[data["SRC"]] > MAX_target_data:
            MAX_target_data = target_data[data["SRC"]]
            MAX_target = data["SRC"]
    logs.append(data)


print(json.dumps(logs, indent=4))


print("Which IP is attacking us? (you have 5 seconds):")

print("[+] DEBUG: " + MAX_target)
def handler(a, b):
    print('You took too long')
    exit(0)


signal.signal(signal.SIGALRM, handler)
signal.alarm(10)

answer = input()

signal.alarm(0) # Disable alarm

if answer.strip() == MAX_target:
    try:
        fh = open("/flag", "r")
    except FileNotFoundError:
        print("FLAG{DEBUG_FLAG}")
        exit(1)
    print(fh.read())
    fh.close()
    exit(0)

print(f"Wrong... It was {MAX_target}")
