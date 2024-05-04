# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22.0
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION} AS build
WORKDIR /src

# Download dependencies as a separate step to take advantage of Docker's caching.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# TODO: mount better than copy
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/ \
    go mod download -x

# This is the architecture you’re building for, which is passed in by the builder.
# Placing it here allows the previous steps to be cached across architectures.
# HACK: Really usefull
ARG TARGETARCH=amd64

# Build the application.
# Leverage a cache mount to /go/pkg/mod/ to speed up subsequent builds.
# TODO: need always use like that in order to not copy files to container 
# HACK: --mount=type=bind,target=. 
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod/ \
    CGO_ENABLED=0 GOARCH=$TARGETARCH go build -o /bin/server ./cmd/proxy

FROM alpine:latest AS development

# Install any runtime dependencies that are needed to run your application.
# Leverage a cache mount to /var/cache/apk/ to speed up subsequent builds.
RUN --mount=type=cache,target=/var/cache/apk \
    apk --update add \
        ca-certificates \
        tzdata \
        && \
        update-ca-certificates

# Create a non-privileged user that the app will run under.
# See https://docs.docker.com/go/dockerfile-user-best-practices/
ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

# Copy the executable from the "build" stage.
COPY --from=build /bin/server /bin/
COPY --from=build --chown=appuser:appuser /src/secrets /secrets

# Just for information
# Expose the port that the application listens on.
EXPOSE 8123

# What the container should run when it is started.
ENTRYPOINT [ "/bin/server" ]
