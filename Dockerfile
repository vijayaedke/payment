
FROM golang:1.22-alpine AS builder

WORKDIR /app

ENV GOPROXY=direct
ENV GODEBUG=netdns=go
RUN apk add --no-cache git

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]