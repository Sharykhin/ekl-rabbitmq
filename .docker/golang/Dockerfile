FROM golang:1.12

WORKDIR /go/src/github.com/Sharykhin/app

COPY main.go .

RUN go build -o main .

CMD ["./main"]