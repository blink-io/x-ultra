#!/bin/bash

TRANS_PATH=github.com/blink-io/x/kratos/v2/tran
protoc --proto_path=. \
  --go-http_out=extern_template=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-go-http \
  ./metadata.proto
