FROM alpine:latest

RUN apk add --update ca-certificates git && rm -rf /var/cache/apk/*
ADD contentd-linux-amd64 /opt/otsimo/content

CMD ["/opt/otsimo/content"]
