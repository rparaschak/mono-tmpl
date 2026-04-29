COMPOSE_FILE := infrastructure/local-dev/docker-compose.yml
COMPOSE_CMD := docker compose -f $(COMPOSE_FILE)

.PHONY: local-env

local-env:
	$(COMPOSE_CMD) up -d
