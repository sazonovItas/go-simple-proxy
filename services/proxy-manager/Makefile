.PHONY: build run 
build:
	go build -o ./bin/app ./cmd/app

run: build
	./bin/app

.PHONY: docker-image docker-container
docker-image:
	@if [ ! -z $(docker images -q proxy-manager-service:latest) ]; then docker image rm $(docker images -q proxy-manager-service:latest); fi;
	@echo "Building docker image proxy-manager-service:latest"
	@docker build --tag=proxy-manager-service:latest .

docker-container:
	docker run --rm -t -i \
		--name=proxy-manager-service \
		--network=host \
		--env-file=./.env.sample \
		-v "./configs/dev.yml:/configs/dev.yml" \
		-v "/var/run/docker.sock:/var/run/docker.sock:rw" \
		-d proxy-manager-service:latest

.PHONY: test
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "Running unit tests with coverage profile"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count
	@go tool cover -func=coverage.txt

.PHONY: clean
clean:
	rm ./bin
