#!/bin/bash


TRANS_PATH=github.com/blink-io/x/kratos/v2/transport/http
EXTERN_TMPL=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl
THIRD_PARTY_PROTOS=/data/projects/open/x/third_party/

DEBUG=1 protoc --proto_path=. \
  --proto_path=$THIRD_PARTY_PROTOS \
  --x-go-http_out=extern_template=${EXTERN_TMPL},transport_path=${TRANS_PATH},paths=source_relative:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-x-go-http \
  ./metadata.proto