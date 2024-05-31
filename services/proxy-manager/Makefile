IMAGE_TAG=latest
SERVICE_NAME=proxy-manager-service

.PHONY: build run 
build:
	go build -o ./bin/$(SERVICE_NAME) ./cmd/$(SERVICE_NAME)

run: build
	./bin/app

.PHONY: docker-image docker-container
docker-image:
	@if [ ! -z $(docker images -q $(SERVICE_NAME):$(IMAGE_TAG)) ]; then docker image rm $(docker images -q $(SERVICE_NAME):$(IMAGE_TAG)); fi;
	@echo "Building docker image $(SERVICE_NAME):$(IMAGE_TAG)"
	@docker build --tag=$(SERVICE_NAME):$(IMAGE_TAG) .

docker-container:
	docker run --rm -t -i \
		--name=$(SERVICE_NAME) \
		--network=host \
		--env-file=./.env.sample \
		-v "./configs/dev.yml:/configs/dev.yml" \
		-v "/var/run/docker.sock:/var/run/docker.sock:rw" \
		-d $(SERVICE_NAME):$(IMAGE_TAG)

.PHONY: test
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo "Running unit tests with coverage profile"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count
	@go tool cover -func=coverage.txt

.PHONY: clean
clean:
	rm ./bin/*
