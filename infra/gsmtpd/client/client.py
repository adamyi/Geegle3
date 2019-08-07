import socket

IP = '127.0.0.1'
PORT = 587

def open_connection(ip, port):
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)

    try:
        s.connect((ip, port))
    except:
        print(f"Can't open connection to server @ {ip}:{port}")
        exit() 
    return s

def recv_until(sock, until):
    message = b''
    new_data = sock.recv(1)    
    while len(new_data):
        message += new_data
        if message.endswith(until):
            return message

        new_data = sock.recv(1)

    print("\n\nConnection closed: \n")
    print(message.decode("UTF-8"))
    exit()

def recv_all(sock):
    return recv_until(sock, b"<--\n")[:-4]

def send(sock, data):
    data += b"\n"
    if sock.send(data) != len(data):
        print("\n\nConnection closed: \n")
        print(sock.recv(8192).decode("UTF-8"))
        exit()

def send_mail(sock):
    to = input("Who do you want to send mail to? ").encode()
    from_ = input("What is your email? ").encode()
    mail = input("What do you want to send?\n").encode()
    
    print("Okay... attempting to send mail...")

    send(sock, b"s")

    recv_until(sock, b"mailto: ")
    send(sock, to)

    recv_until(sock, b"mailfrom: ")
    send(sock, from_)

    recv_until(sock, b"mail: ")
    send(sock, mail)
    
    print(recv_all(sock).decode('UTF-8'))

def get_mail(sock):
    email = input("What is your email? ").encode()
    
    print("Okay... attempting to read mail...")

    send(sock, b"g")

    recv_until(sock, b"email: ")
    send(sock, email)

    print(recv_all(sock).decode('UTF-8'))

def main():
    print("Attempting to connect to mailserver...")
    sock = open_connection(IP, PORT)
       
    while True:
        print("Loading menu...")
        recv_until(sock, b"> ") # Wait until connected
        
        option = input("Would you like to [r]ead mail or [s]end mail or [q]uit: ")

        if option == 'r':
            get_mail(sock)
        elif option == 's':
            send_mail(sock)
        elif option == 'q':
            break
        else:
            print("Incorrect option selected...")

    sock.close()

if __name__ == "__main__":
    main()

