#!/bin/sh
dir=$(pwd)
echo "Switching to GOPATH"
cd "$GOPATH/src"
echo "Generating protobuf code"
protoc --go_out=plugins=grpc:. github.com/bio-routing/tflow3/pkg/flow/*.proto
echo "Switching back to working directory"
cd $dir

