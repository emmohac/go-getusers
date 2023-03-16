FROM golang:1.20

RUN mkdir -p /app

COPY . /app/

WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]