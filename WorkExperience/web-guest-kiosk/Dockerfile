FROM golang:1.11

RUN apt-get update
RUN apt-get install -y --no-install-recommends \
    libzbar-dev

RUN go get "github.com/clsung/grcode"
RUN go get "github.com/gorilla/mux"
RUN go get "github.com/skip2/go-qrcode"
RUN go get "golang.org/x/crypto/acme/autocert"

COPY app /kiosk

WORKDIR /kiosk
CMD go run main.go helper.go 
