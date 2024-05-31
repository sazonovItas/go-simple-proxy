ENV_FILE=./.env/.env.dev

.PHONY: lint
lint:
	./scripts/lint.sh

.PHONY: clean_none_images
clean_none:
	./scripts/clean_none_images.sh

.PHONY: docker-up docker-down
docker-up:
	docker compose --env-file $(ENV_FILE) up -d

docker-down:
	docker compose --env-file $(ENV_FILE) down

