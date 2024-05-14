
build:
	go build -o ./.bin/proxy ./cmd/proxy
	go build -o ./.bin/proxy_manager ./cmd/proxy_manager
.PHONY: build

run: build
	DOCKER_API_VERSION=1.44 ./.bin/proxy_manager
.PHONY: run

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -race -v ./...
.PHONY: test

proxy-test:
	curl --proxytunnel -v --proxy http://0.0.0.0:8123 --proxy-insecure -k https://mangalib.me
.PHONY: proxy-test

coverage:
	go test -coverprofile=c.out ./...;\
	go tool cover -func=c.out;\
	rm c.out
.PHONY: coverage
