#!/bin/bash

projectPath=$(pwd)

cd ../../

export GOPATH=$(pwd)

echo "GOPATH=$GOPATH"

cd ${projectPath}

echo "dep ensure"

dep ensure -v

if [ $? -eq 0 ]; then
  echo "build..."
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o peckergo ./api
  chmod +x peckergo
else
  echo "dep ensure fail"
  exit 1
fi





