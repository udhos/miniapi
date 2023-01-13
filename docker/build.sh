#!/bin/bash

version=$(go run ./cmd/miniapi -version | awk '{ print $2 }' | awk -F= '{ print $2 }')

echo version=$version

docker build \
    -t udhos/miniapi:latest \
    -t udhos/miniapi:$version \
    -f docker/Dockerfile .