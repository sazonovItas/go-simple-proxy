#!/usr/bin/env sh

# TODO: use docker images filter
docker rmi -f $(docker images | grep none | awk '{print $3}')
