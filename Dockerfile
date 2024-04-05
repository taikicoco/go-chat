FROM golang:1.20-alpine

ENV CGO_ENABLED=1

RUN apk update
RUN apk upgrade
RUN apk add --update figlet
RUN apk add build-base

RUN go install github.com/99designs/gqlgen@latest
