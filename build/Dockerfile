FROM golang:1.19.1-alpine3.15
RUN apk add --update && apk add git
# Make app directory
RUN mkdir /go/src/app
# Set working directory
WORKDIR /go/src/app
# Move host files to working directory in container
ADD . /go/src/app