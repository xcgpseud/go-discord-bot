FROM golang:1.13.6-alpine3.11

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]