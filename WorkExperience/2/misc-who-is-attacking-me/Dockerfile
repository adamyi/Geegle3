FROM ubuntu:xenial

# Apt-get update
RUN apt-get update

# Install xinetd
RUN apt-get install -y xinetd python3
RUN update-rc.d xinetd disable

# Add the ctf dir
RUN mkdir /home/ctf

RUN chmod 775 /home/ctf

# Add the ctf user
RUN useradd -M -U\
    -d /home/ctf \
    ctf

RUN chown ctf:ctf /home/ctf

# Add the content
COPY chal/gen.py /home/ctf/pwn
COPY xinetd.conf /etc/xinetd.conf
COPY flag /flag

# Make the binary
WORKDIR /home/ctf

# Expose the service port
EXPOSE 1000

# Clean up setup files
RUN ["chmod","775","/flag"]
RUN ["chmod","775","/home/ctf/pwn"]
CMD ["script","-c", "xinetd -d -dontfork","/dev/null"]
