#!/usr/bin/env sh

for file in $(find . -name go.mod); do
  echo "============ Linting $(dirname $file) ============="
  golangci-lint run --config .golangci.yml "$(dirname $file)/..."
done;
