.PHONY: build run generate-keys proxy-test

generate-keys:
	cd ./secrets && ../scripts/generate_ssl_keys.sh

proxy-test:
	curl --proxytunnel -v --proxy http://0.0.0.0:8123 --proxy-insecure -k https://mangalib.me

build:
	go build -o ./bin/go-simple-proxy ./cmd/proxy/main.go

run: build
	./bin/go-simple-proxy

