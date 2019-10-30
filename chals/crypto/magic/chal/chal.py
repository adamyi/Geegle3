#!/usr/bin/env python3

def encrypt(data, key):
    if len(data) != len(key):
        print("Length didn't match")
        exit(0)

    ans = ''
    for i in range(len(key)):
        ans += chr(ord(key[i]) ^ ord(data[i]))

    return ans


fh = open("/flag", "r")
flag = fh.read()


print("Time to roll my own crypto")
print("Enter some data and I'll encrypt it with the flag as a key")
print("Because this is in testing, the data you enter has to be exactly " + str(len(flag)) + " characters long")

print("Encrypted data is " + encrypt(str(input()), str(flag)))
