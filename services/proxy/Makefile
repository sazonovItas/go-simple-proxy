.PHONY: build run 
build:
	go build -o ./bin/app ./cmd/app

run: build
	./bin/app

.PHONY: docker-image docker-container
docker-image:
	@if [ ! -z $(docker images -q go-proxy:latest) ]; then docker image rm $(docker images -q go-proxy:latest); fi;
	@echo "Building docker image go-proxy:latest"
	@docker build --tag=go-proxy:latest .

docker-container:
	docker run --rm -t -i \
		--name=go-proxy-test \
		--network=host \
		--env-file=./.env.sample \
		-d go-proxy:latest

.PHONY: test
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "Running unit tests with coverage profile"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count
	@go tool cover -func=coverage.txt

.PHONY: clean
clean:
	rm ./bin
