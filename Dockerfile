FROM docker.io/library/busybox:latest
MAINTAINER lgh
WORKDIR /tmp
COPY main /tmp
COPY /configs /tmp/configs
COPY config /root/.kube/config
COPY /pkg/templates /tmp/pkg/templates
COPY /pkg/static /tmp/pkg/static
EXPOSE 8000
CMD ./main
