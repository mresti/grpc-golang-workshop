FROM golang:1.10.3-alpine

ENV SRV_NAME step-7-deploy-grpc-server

ENV PKG_PATH /go/src/github.com/mresti/grpc-golang-workshop

ADD books $PKG_PATH/$SRV_NAME/books
ADD server.go $PKG_PATH/$SRV_NAME/server.go

WORKDIR $PKG_PATH/$SRV_NAME

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
RUN go get -u google.golang.org/grpc golang.org/x/net/context github.com/golang/protobuf/proto github.com/olebedev/emitter

RUN go install

EXPOSE 50051
