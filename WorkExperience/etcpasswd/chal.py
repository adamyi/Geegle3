import os

print("Welcome to sysadmin control system")
username = input("Username: ")
password = input("Password: ")

if username == "root" and password == "cocacola":
    os.system("cat flag")
