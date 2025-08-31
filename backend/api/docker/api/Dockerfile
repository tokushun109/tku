FROM golang:1.25-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY ./ ${ROOT}

RUN apk update && apk add git && go mod download
