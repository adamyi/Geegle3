FROM golang:1.11

WORKDIR $GOPATH/src/adamyi/geegle-sffe
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["sffe"]
