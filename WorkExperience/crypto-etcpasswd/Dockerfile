FROM python:latest

RUN apt-get update

RUN apt-get install -y wget unzip

COPY chal /app
COPY flag /flag
COPY static /static

WORKDIR /app

RUN useradd -M -U\
    -d /app \
    ctf

RUN chown ctf:ctf /app

RUN wget https://github.com/joewalnes/websocketd/releases/download/v0.3.0/websocketd-0.3.0-linux_amd64.zip && unzip websocketd-0.3.0-linux_amd64.zip

RUN ["chmod","775","/flag"]
RUN ["chmod","775","/app/app.py"]
CMD ["./websocketd", "--port=80", "--staticdir=/static", "--binary=true", "/app/app.py"]
