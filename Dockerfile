FROM golang:1.14-alpine

WORKDIR /app

COPY . /app

RUN apk add xargs

RUN go build -o main

RUN ls | sed 's/^main//' | xargs -L1 rm -rf

ENTRYPOINT ["./main"]