FROM golang:1.20

WORKDIR /go/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["air"]