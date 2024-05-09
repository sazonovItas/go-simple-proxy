generate-keys:
	cd ./secrets && ../scripts/generate_ssl_keys.sh
.PHONY: generate-keys

proxy-test:
	curl --proxytunnel -v --proxy http://0.0.0.0:8123 --proxy-insecure -k https://mangalib.me
.PHONY: proxy-test

build:
	go build -o ./.bin/go-simple-proxy ./cmd/proxy/main.go
.PHONY: build

run: build
	./.bin/go-simple-proxy
.PHONY: run

lint:
	golangci-lint run ./...
.PHONY: lint

test:
	go test -race -v ./...
.PHONY: test

coverage:
	go test -coverprofile=c.out ./...;\
	go tool cover -func=c.out;\
	rm c.out
.PHONY: coverage
