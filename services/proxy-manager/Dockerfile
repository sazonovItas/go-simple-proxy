# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.2
FROM golang:${GO_VERSION} AS build
WORKDIR /src

RUN --mount=type=cache,target=/go/pkg/mod/  \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download -x

ARG TARGETARCH=amd64
RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=.             \
    CGO_ENABLED=0 GOARCH=${TARGETARCH} go build -o /bin/app ./cmd/app

FROM alpine:3.19.1 AS development

COPY --from=build /bin/app /bin/

EXPOSE 3030

ENTRYPOINT [ "/bin/app" ]
