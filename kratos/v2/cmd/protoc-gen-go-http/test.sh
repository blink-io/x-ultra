#!/bin/bash

PARAMS="paths=source_relative:.,extern_template=/data/projects/open/x/kratos/v2/cmd/protoc-gen-go-http/httpTemplate-Blink.tpl"

protoc --proto_path=. \
  --go-http_out=$PARAMS \
  i18n.proto
