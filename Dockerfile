FROM golang:1.7-alpine

WORKDIR /go/src/github.com/elmariofredo/tc-agent-name-unlocker

RUN apk update && \
    apk add --no-cache git zsh

RUN go get -v github.com/spf13/cobra/cobra

ENTRYPOINT [ "zsh" ]
