#!/bin/bash

protoc --proto_path=. \
  --go-http_out=extern_template=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-go-http \
  ./i18n.proto
