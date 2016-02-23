FROM alpine:3.3
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

RUN apk add --update git && rm -rf /var/cache/apk/*

ADD bin/otsimo-content-linux-amd64 /opt/otsimo-content/bin/otsimo-content

EXPOSE 18859

CMD ["/opt/otsimo-content/bin/otsimo-content"]
