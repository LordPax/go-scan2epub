#!/bin/bash

[ -z "$1" ] && echo "Usage: $0 <version>" && exit 1

echo "Building version $1"

docker build \
    --platform linux/arm64,linux/amd64 \
    --build-arg VERSION=$1 \
    -t lordpax/scan2epub:$1 \
    -t lordpax/scan2epub:latest \
    --push .

[ $? -eq 0 ] && echo "Build successful" || echo "Build failed"
