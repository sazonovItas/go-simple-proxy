#!/usr/bin/env sh

set -e

if [ ! -z "$(docker images -q go-proxy > /dev/null)" ]; then
  docker image rm $(docker images -q go-proxy)
fi

docker build -f ./docker/proxy.Dockerfile --tag go-proxy .
