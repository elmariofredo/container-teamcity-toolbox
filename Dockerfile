FROM alpine:3.4

COPY ./container-teamcity-toolbox /usr/local/bin/container-teamcity-toolbox

COPY ./tc-agent-names-cleanup.sh /etc/periodic/15min/tc-agent-names-cleanup

ENTRYPOINT [ "crond", "-f" ]
