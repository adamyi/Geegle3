FROM golang:1.11

WORKDIR $GOPATH/src/geegle.org/geemail-frontend
COPY main.go .
COPY dist/geemail-app /files
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["geemail-frontend"]
