#!/usr/bin/env sh

set -e

if [ ! -z "$(docker images -q go-proxy > /dev/null)" ]; then
  docker image rm $(docker images -q go-proxy)
fi

docker build -f ./services/proxy/Dockerfile --tag go-proxy ./services/proxy

if [ ! -z "$(docker images -q go-proxy-manager > /dev/null)" ]; then
  docker image rm $(docker images -q go-proxy-manager)
fi

docker build -f ./services/proxy-manager/Dockerfile --tag go-proxy-manger ./services/proxy-manager

docker compose up -d 
