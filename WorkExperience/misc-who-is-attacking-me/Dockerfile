FROM ubuntu:xenial

RUN apt-get update

RUN apt-get install -y python3 wget unzip

RUN mkdir /home/ctf

RUN chmod 775 /home/ctf

RUN useradd -M -U\
    -d /home/ctf \
    ctf

RUN chown ctf:ctf /home/ctf

COPY chal/gen.py /home/ctf/pwn
COPY static /home/ctf/static
COPY flag /flag

WORKDIR /home/ctf

RUN wget https://github.com/joewalnes/websocketd/releases/download/v0.3.0/websocketd-0.3.0-linux_amd64.zip && unzip websocketd-0.3.0-linux_amd64.zip

RUN ["chmod","775","/flag"]
RUN ["chmod","775","/home/ctf/pwn"]
CMD ["./websocketd", "--port=80", "--staticdir=/home/ctf/static", "--binary=true", "/home/ctf/pwn"]
