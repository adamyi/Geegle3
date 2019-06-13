# Use an official Python runtime as an image
FROM python:2.7

COPY app /gae

WORKDIR /gae
RUN pip install -r requirements.txt

CMD gunicorn main:app -b 0.0.0.0:80
