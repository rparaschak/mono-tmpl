dev:
	go run github.com/air-verse/air@latest

migrate:
	atlas migrate diff $(NAME) --env $(ENV)
	atlas migrate apply --env $(ENV)
