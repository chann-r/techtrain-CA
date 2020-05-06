FROM golang:1.14-alpine

WORKDIR /techtrain-CA

ENV GO111MODULE=on

COPY go.mod .
# COPY go.sum .

RUN go mod download

EXPOSE 8080
