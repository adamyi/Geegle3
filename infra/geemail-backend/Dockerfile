FROM golang:1.11

WORKDIR $GOPATH/src/geegle.org/geemail-backend
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["geemail-backend"]
