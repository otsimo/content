FROM centurylink/ca-certs
MAINTAINER Sercan Degirmenci <sercan@otsimo.com>

ADD bin/otsimo-content-linux-amd64 /opt/otsimo-content/bin/otsimo-content

EXPOSE 18859

CMD ["/opt/otsimo-content/bin/otsimo-catalog"]
