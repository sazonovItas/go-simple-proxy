# syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM alpine:3.19.1
RUN apk add --update docker openrc && rc-update add docker boot
