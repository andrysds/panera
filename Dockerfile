FROM alpine:latest

MAINTAINER Andrys <andrysds@gmail.com>

WORKDIR "/opt"

ADD .docker_build/pnr_bot /opt/bin/pnr_bot

CMD ["/opt/bin/pnr_bot"]
