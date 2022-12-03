FROM --platform=linux/amd64 node:16-alpine

ENV ROOT=/client
WORKDIR ${ROOT}
COPY ./ ${ROOT}

EXPOSE 3000
ENV HOST 0.0.0.0

RUN apk update && \
    apk add git && \
    yarn install --production=true && \
    yarn build