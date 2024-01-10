#!/bin/bash

THIRD_PARTY_PROTOS=/data/projects/open/x/third_party/

DEBUG=1 protoc --proto_path=. \
  --proto_path=$THIRD_PARTY_PROTOS \
  --x-connect-go_out=paths=source_relative:. \
  --plugin /home/heisonyee/go/bin/protoc-gen-x-connect-go \
  ./metadata.proto