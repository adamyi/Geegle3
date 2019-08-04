FROM python:latest

RUN apt-get update

# Install xinetd
RUN apt-get install -y xinetd
RUN update-rc.d xinetd disable

# Add the content
COPY chal /app
COPY flag /flag
COPY xinetd.conf /etc/xinetd.conf

WORKDIR /app

# add the ctf user
RUN useradd -M -U\
    -d /app \
    ctf

RUN chown ctf:ctf /app
EXPOSE 1000

# Clean up setup files
CMD ["chmod", "777", "/app/app.py"]
CMD ["chmod", "777", "/flag"]
CMD ["script", "-c", "xinetd -d -dontfork","/dev/null"]
