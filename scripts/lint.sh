#!/usr/bin/env sh

golangci-lint run --config .golangci.yml $(find . -type f -name go.mod | sed -r 's|/[^/]+$||' | sed -e 's/$/\/.../')
