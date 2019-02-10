FROM golang:1.9.0-alpine

ENV CGO_ENABLED=0 GOOS=linux
ENV GIT_TERMINAL_PROMPT=1

WORKDIR /go/src/github.com/synergydesigns/stylesblitz-server

COPY . /go/src/github.com/synergydesigns/stylesblitz-server

RUN apk add --no-cache curl \
    make \
    bash \
    git \
    nodejs \
    yarn  \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && yarn global add supervisor

ENV ROOT_DIRECTORY=/go/src/github.com/synergydesigns/stylesblitz-server
