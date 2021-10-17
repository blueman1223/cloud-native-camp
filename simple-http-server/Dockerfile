# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY *.go ./
RUN go mod init main
RUN go build -o /simple-http-server

ENV MY_SERVICE_PORT=8888

EXPOSE 8888

ENTRYPOINT ["/simple-http-server"]

