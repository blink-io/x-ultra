#!/bin/bash


#TRANS_PATH=github.com/blink-io/x/kratos/v2/transport/http
#EXTERN_TMPL=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl
THIRD_PARTY_PROTOS=/data/projects/open/x/third_party/

protoc --proto_path=. \
  --proto_path=$THIRD_PARTY_PROTOS \
  --go_out=paths=source_relative:. \
  --blink-go-http_out=paths=source_relative:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-blink-go-http \
  ./metadata.proto