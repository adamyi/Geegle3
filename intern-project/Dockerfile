# Basic Xinetd Docker Container
# Useful for anyting binary / requiring netcat

FROM ubuntu:xenial

# Add 32 bit support
RUN dpkg --add-architecture i386

# Apt-get update
RUN apt-get update

# Add 32 bit libraries
RUN apt-get install -y libc6:i386 libstdc++6:i386

# Install xinetd
RUN apt-get install -y xinetd
RUN update-rc.d xinetd disable

# Install nc for debugging
RUN apt-get install -y netcat

# Add the ctf dir
RUN mkdir /home/ctf

RUN chmod 775 /home/ctf

# Add the ctf user
RUN useradd -M -U\
    -d /home/ctf \
    ctf

RUN chown ctf:ctf /home/ctf

# Add the content
COPY src/intern-project /home/ctf/pwn
COPY xinetd.conf /etc/xinetd.conf
COPY flag /flag

RUN chmod 775 /var/mail

# Make the binary
WORKDIR /home/ctf

# Expose the service port
EXPOSE 19204

# Clean up setup files
RUN ["chmod","775","/flag"]
RUN ["chmod","775","/home/ctf/pwn"]
CMD ["script","-c", "xinetd -d -dontfork","/dev/null"]
