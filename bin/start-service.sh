#!/usr/bin/env bash

docker service create --name teamcity-agent-lock-cleanup \
 --mount type=bind,src=/tmp,dst=/opt/docker-shared \
 --mount type=bind,src=/var/run,dst=/var/run \
 --mode global \
container-teamcity-toolbox:latest