#!/usr/bin/env bash

docker build -f Dockerfile.remove-lock -t container-teamcity-toolbox:remove-lock-$1 .