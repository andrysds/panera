FROM alpine:latest

MAINTAINER Andrys <andrysds@gmail.com>

WORKDIR "/opt"

ADD .docker_build/panera /opt/bin/panera

CMD ["/opt/bin/panera"]
