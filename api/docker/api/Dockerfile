FROM golang:1.16.2-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY ./ ${ROOT}

RUN apk update && apk add git
RUN go mod download