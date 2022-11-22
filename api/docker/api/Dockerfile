FROM golang:1.18-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY ./ ${ROOT}

RUN apk update && apk add git && \
    go install github.com/cosmtrek/air@latest && \
    go mod download