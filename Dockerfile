FROM golang:1.20-alpine

RUN apk update && apk add --no-cache git
RUN apk add --no-cache bash

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o /binary 

EXPOSE 8080

ENTRYPOINT [ "/binary" ]