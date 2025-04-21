include .env
export

DATABASE_URL = mysql://$(DB_USER):$(DB_PASSWORD)@localhost:$(DB_PORT)/$(DB_NAME)

.PHONY: schemadiff
schemadiff: internal/schema
	atlas migrate diff --env gorm

.PHONY: migrate
migrate: schemadiff
	atlas migrate apply --url $(DATABASE_URL)

.PHONY: initial_migrate
initial_migrate:
	atlas migrate apply --url $(DATABASE_URL)

.PHONY: swag_init
swag_init: swag_fmt
	swag init -g cmd/api/main.go --output cmd/api/docs/swagger

.PHONY: swag_fmt
swag_fmt: cmd/api internal/interface
	swag fmt ./

.PHONY: schemaspy
schemaspy: internal/schema
	docker compose up schemaspy