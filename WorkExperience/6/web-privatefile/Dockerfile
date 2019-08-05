FROM golang:1.11

RUN apt-get update

RUN go get "github.com/gorilla/mux"

COPY app /filebrowser

WORKDIR /filebrowser
CMD go run main.go helper.go 
