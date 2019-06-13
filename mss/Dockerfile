FROM golang:1.11

WORKDIR $GOPATH/src/geegle.org/mss
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["mss"]
