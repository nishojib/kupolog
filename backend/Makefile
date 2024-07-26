ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## audit: tidy and vendor dependencies and format, vet and test all code
.PHONY: audit
audit: vendor
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...

## vendor: tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# BUILD
# ==================================================================================== #
## install: install dependencies
.PHONY: install
install: 
	go get ./...
	go mod tidy
	go mod download


## build/api: build the cmd/api application
.PHONY: build/api
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s -w' -o=/bin/api ./cmd/api

# ==================================================================================== #
# RUN
# ==================================================================================== #
## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	@echo 'Running cmd/api...'
	go run ./cmd/api


# ==================================================================================== #
# SWAGGER
# ==================================================================================== #
## swagger: generate the swagger documentation
.PHONY: swagger
swagger:
	@echo 'Generating swagger documentation...'
	swag init -g ./cmd/api/main.go

# ==================================================================================== #
# DATABASE
# ==================================================================================== #
## db/create: create a new database migration
.PHONY: db/create
db/create:
	goose -s -dir="migrations" create ${name} sql

## db/up: migrate the database
.PHONY: db/up
db/up:
	goose -dir="migrations" turso ${DB_DSN} up

## db/down: rollback the database
.PHONY: db/down
db/down:
	goose -dir="migrations" turso ${DB_DSN} down