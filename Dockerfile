FROM golang:1.12.7-alpine

ENV CGO_ENABLED=0 GOOS=linux
ENV GIT_TERMINAL_PROMPT=1
ENV GO111MODULE=on

WORKDIR /go/src/github.com/synergydesigns/stylesblitz-server
COPY ./ .

RUN apk add --no-cache curl \
    make \
    bash \
    git \
    nodejs \
    yarn  \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && yarn global add supervisor \
    && yarn global add serverless

ENV ROOT_DIRECTORY=/go/src/github.com/synergydesigns/stylesblitz-server
