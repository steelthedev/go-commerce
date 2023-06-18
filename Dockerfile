FROM golang:1.18.0

WORKDIR /app


RUN go install github.com/cosmtrek/air@latest

COPY . /app

RUN go mod tidy