#!/usr/bin/env sh

docker rmi -f $(docker images | grep none | awk '{print $3}')
