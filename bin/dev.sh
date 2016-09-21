#!/bin/bash

docker run -v /var/run:/var/run \
    --volume=$PWD:/go/src/github.com/elmariofredo/tc-agent-name-unlocker \
    --volume=$PWD/tmp/agentlock:/opt/docker-shared/agentlock \
    --interactive \
    --tty \
    tc-agent-name-unlocker:16c9f367cc2f926c0a8d626b7c4ebbc714c721ec \
    -c "go run main.go cleanup"