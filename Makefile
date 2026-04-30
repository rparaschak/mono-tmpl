COMPOSE_FILE := infrastructure/local-dev/docker-compose.yml
COMPOSE_CMD := docker compose -f $(COMPOSE_FILE)

.PHONY: local-env
.PHONY: openapi

local-env:
	$(COMPOSE_CMD) up -d

openapi:
	go -C api run ./cmd/openapi -out openapi.json
