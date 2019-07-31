FROM golang:1.11

RUN apt-get update

RUN go get "github.com/gorilla/mux"

COPY app /easy_ssrf

WORKDIR /easy_ssrf
CMD go run main.go helper.go 
