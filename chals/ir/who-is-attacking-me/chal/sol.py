#!/usr/bin/env python2
import json
from pwn import *

p = process("python cli-relay.py eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiY2RAZ2VlZ2xlLm9yZyIsInNlcnZpY2UiOiJ1YmVycHJveHlAc2VydmljZXMuZ2VlZ2xlLm9yZyIsImV4cCI6MTU2NzczODg5MX0.XGvQlcCcfwb6TYc-UrhMqDLRzC_R0JkXGAWIEqju5ys attack-check.corp.geegle.org", shell=True)
#p = process('./gen.py')

p.recvuntil("data to ")
target = p.recvline().strip()
log.info("Target is " + target)

p.recvuntil("Ready?\n")
data = json.loads(p.recvuntil("\nWhich")[:-6])

MAX_target = "empty"
MAX_target_data = 0
target_data = defaultdict(int)

for d in data:
    if d["DST"] == target:
        target_data[d["SRC"]] = target_data[d["SRC"]] + int(d["data_size"])

        if target_data[d["SRC"]] > MAX_target_data:
            MAX_target = d["SRC"]
            MAX_target_data = target_data[d["SRC"]]

p.sendline(MAX_target)
p.interactive()

