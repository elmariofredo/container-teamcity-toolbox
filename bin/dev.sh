#!/usr/bin/env bash

docker run -v /var/run:/var/run \
    --volume=$PWD:/go/src/github.com/elmariofredo/container-teamcity-toolbox \
    --volume=$PWD/tmp/agentlock:/opt/docker-shared/agentlock \
    --interactive \
    --tty \
    container-teamcity-toolbox:dev \
    -c "$1"