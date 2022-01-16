FROM golang:1.17.5-alpine3.15 AS build

WORKDIR /opt/document-server

COPY ./bin/document-server /opt/document-server
COPY docker/dev/entrypoint.sh  /opt/document-server/entrypoint.sh

# set Time Zone
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV PATH "$PATH:/root/go/bin"
ENV GOPATH "/root/go/src"
ENV APP_ENV test

RUN chmod +x /opt/entrypoint.sh
RUN chmod +x document-server

ENTRYPOINT /opt/entrypoint.sh