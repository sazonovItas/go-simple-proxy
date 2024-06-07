IMAGE_TAG=latest
SERVICE_NAME=auth-service

.PHONY: build run 
build:
	go build -o ./bin/$(SERVICE_NAME) ./cmd/app

run: build
	DATABASE_URI='postgres://proxy_manager_user:proxy_manager_password@127.0.0.1:5432/proxy_manager_db?sslmode=disable' ./bin/$(SERVICE_NAME)

.PHONY: proto
proto:
	$(foreach proto_file, $(shell find api/proto -name '*.proto'),\
		protoc --go_out=pkg/pb/ --go-grpc_out=pkg/pb/ \
		--proto_path=api/proto  $(proto_file);)

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
