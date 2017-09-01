FROM golang:alpine 
LABEL maintainer "Darwin Smith II <dwin@dlsmi.com>"

LABEL app_version="0.1.0" architecture="amd64"

WORKDIR /go/src/app
COPY /app /go/src/app

RUN go-wrapper get
RUN go-wrapper install

EXPOSE 8080
CMD ["go-wrapper", "run"] # ["app"]