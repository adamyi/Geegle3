# Use an official Python runtime as an image
FROM python:3.7

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    pandoc \
    texlive-latex-base \
    lmodern \
    texlive-fonts-recommended

COPY app /kix
COPY key /key

WORKDIR /kix
RUN pip install -r requirements.txt

CMD gunicorn main:app -b 0.0.0.0:80
