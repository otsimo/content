FROM alpine:3.5

RUN apk add --update ca-certificates git && rm -rf /var/cache/apk/*
ADD otsimo-content-linux-amd64 /opt/otsimo/content

CMD ["/opt/otsimo/content"]
