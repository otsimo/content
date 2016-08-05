FROM alpine:3.4
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

RUN apk add --update ca-certificates git && rm -rf /var/cache/apk/*

ADD otsimo-content-linux-amd64 /opt/otsimo/content

EXPOSE 18859

CMD ["/opt/otsimo/content"]
