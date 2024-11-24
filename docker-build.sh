#!/bin/bash

[ -z "$1" ] && echo "Usage: $0 <version>" && exit 1
echo "Building version $1"
docker build --build-arg VERSION=$1 -t lordpax/scan2epub:$1 -t lordpax/scan2epub:latest .
docker push lordpax/scan2epub:$1
docker push lordpax/scan2epub:latest
