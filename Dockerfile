FROM golang:1.7-alpine

WORKDIR /go/src/github.com/elmariofredo/tc-agent-name-unlocker

RUN apk update && \
    apk add --no-cache git zsh

RUN go get github.com/spf13/cobra/cobra
RUN go get github.com/docker/docker/client
RUN go get github.com/docker/docker/api/types
RUN go get golang.org/x/net/context 

ENTRYPOINT [ "zsh" ]
