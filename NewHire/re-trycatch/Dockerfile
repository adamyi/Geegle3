# Basic Xinetd Docker Container
# Useful for anyting binary / requiring netcat

FROM ubuntu:xenial

# Add 32 bit support
RUN dpkg --add-architecture i386

# Apt-get update
RUN apt-get update

# Add 32 bit libraries
RUN apt-get install -y libc6:i386 libstdc++6:i386 wget unzip

# Add the ctf dir
RUN mkdir /home/ctf

RUN chmod 775 /home/ctf

# Add the ctf user
RUN useradd -M -U\
    -d /home/ctf \
    ctf

RUN chown ctf:ctf /home/ctf

# Add the content
COPY chal/trycatch /home/ctf/pwn
COPY static /home/ctf/static
COPY flag /flag

# Make the binary
WORKDIR /home/ctf

RUN wget https://github.com/joewalnes/websocketd/releases/download/v0.3.0/websocketd-0.3.0-linux_amd64.zip && unzip websocketd-0.3.0-linux_amd64.zip

RUN ["chmod","775","/flag"]
RUN ["chmod","775","/home/ctf/pwn"]
CMD ["./websocketd", "--port=80", "--staticdir=/home/ctf/static", "/home/ctf/pwn"]
