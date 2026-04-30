.PHONY: build
.PHONY: check
.PHONY: dev
.PHONY: lint
.PHONY: migration
.PHONY: test
.PHONY: vet

TEST_COMPOSE_FILE := ../infrastructure/test/docker-compose.yml
TEST_COMPOSE_PROJECT := mono-tmpl-test
TEST_COMPOSE_CMD := docker compose -p $(TEST_COMPOSE_PROJECT) -f $(TEST_COMPOSE_FILE)
TEST_DATABASE_URL := postgresql://supabase_admin:docker@localhost:5102/mono-tmpl-test?search_path=public&sslmode=disable
GOLANGCI_LINT_VERSION := v2.11.4
GOLANGCI_LINT := go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

check: vet lint build test

dev:
	go run github.com/air-verse/air@latest

vet:
	go vet ./...

lint:
	$(GOLANGCI_LINT) run

build:
	@set -e; \
	build_dir=$$(mktemp -d); \
	cleanup() { rm -rf "$$build_dir"; }; \
	trap cleanup EXIT; \
	go build -o "$$build_dir/" ./...

migration:
	atlas migrate diff $(NAME) --env $(ENV)
	atlas migrate apply --env $(ENV)

test:
	@set -e; \
	log=$$(mktemp); \
	cleanup() { $(TEST_COMPOSE_CMD) down -v --remove-orphans >/dev/null 2>&1 || true; rm -f "$$log"; }; \
	fail() { step="$$1"; printf '\n%s failed:\n' "$$step"; cat "$$log"; exit 1; }; \
	trap cleanup EXIT; \
	: > "$$log"; \
	$(TEST_COMPOSE_CMD) down -v --remove-orphans >"$$log" 2>&1 || fail "docker compose down"; \
	: > "$$log"; \
	$(TEST_COMPOSE_CMD) up -d --wait >"$$log" 2>&1 || fail "docker compose up"; \
	: > "$$log"; \
	atlas migrate apply --dir file://migrations --url '$(TEST_DATABASE_URL)' >"$$log" 2>&1 || fail "atlas migrate apply"; \
	: > "$$log"; \
	APP_ENV=autotest DATABASE_URL='$(TEST_DATABASE_URL)' go test -tags=integration ./... >"$$log" 2>&1 || fail "go test"; \
	printf 'Tests passed.\n'
