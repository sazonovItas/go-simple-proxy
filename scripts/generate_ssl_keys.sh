#!/usr/bin/env sh

openssl genrsa -out server.key 4096
openssl req -new -x509 -nodes -key server.key -out server.crt
