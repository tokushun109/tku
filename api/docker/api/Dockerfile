FROM golang:1.18-alpine

ENV ROOT=/api
WORKDIR ${ROOT}
COPY ./ ${ROOT}

RUN apk update && apk add git
RUN go mod download