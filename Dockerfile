FROM golang:alpine 
LABEL maintainer "Darwin Smith II <dwin@dlsmi.com>"
LABEL app_version="0.1.0" architecture="amd64"
RUN echo $'#!/bin/sh\nntpd -d -q -n -p us.pool.ntp.org' > /etc/periodic/hourly/getntp
RUN apk add --no-cache git
VOLUME /static
VOLUME /log
ENV GIN_MODE=release

WORKDIR /go/src/app
COPY /app /go/src/app

RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080
CMD ["go-wrapper", "run"] # ["app"]
# docker build . -t dwin/gotestserver:0.1.0
# docker run -d -v ~/Docker/gotestserver/log:/log -v ~/Docker/gotestserver/static:/static -p 8080:8080 dwin/gotestserver:0.1.0