FROM golang:1.9.0-alpine

ENV CGO_ENABLED=0 GOOS=linux

WORKDIR /go/src/gitlab.com/synergy-designs/style-blitz

COPY . /go/src/gitlab.com/synergy-designs/style-blitz
ADD . /go/src/gitlab.com/synergy-designs/style-blitz

RUN apk add --no-cache curl \
    make \
    bash \
    git \
    nodejs \
    yarn  \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && /go/bin/dep ensure \
    && yarn global add supervisor

RUN make build && make migrate

ENV ROOT_DIRECTORY=/go/src/gitlab.com/synergy-designs/style-blitz

CMD [ "make", "supervise" ]