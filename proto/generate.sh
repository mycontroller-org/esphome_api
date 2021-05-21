#!/bin/bash

# pre-requests
# go get -u google.golang.org/protobuf/cmd/protoc-gen-go
# go install google.golang.org/protobuf/cmd/protoc-gen-go

OUT_DIR=../pkg/api
mkdir -p ${OUT_DIR}

protoc \
  --go_out=${OUT_DIR} --go_opt=paths=source_relative \
  --go_opt=Mapi.proto=pkg/api \
  --go_opt=Mapi_options.proto=pkg/api \
  *.proto
