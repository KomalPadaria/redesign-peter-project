MIGRATION_FILE = $(shell date +"migrations/%Y%m%d%H%M%S-$(name).sql")

setup-goimports: ## Install the goimports
	go install -mod=mod golang.org/x/tools/cmd/goimports@latest

setup-covmerge: ## Install the covmerge tool
	go get github.com/wadey/gocovmerge
	go install -mod=mod github.com/wadey/gocovmerge

setup-migrate: ## Install the migrate tool
	go install -mod=mod github.com/nurdsoft/redesign-grp-trust-portal-api/shared/db/migrate

setup-mockgen: ## Install mockgen to generate mocks
	go install github.com/golang/mock/mockgen@latest

setup: setup-covmerge setup-goimports setup-migrate setup-mockgen## Install all the build and lint dependencies

dep: ## Get all dependencies
	go env -w GOPROXY=direct
	go env -w GOSUMDB=off
	go mod download
	go mod tidy
	go mod vendor

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done

lint: dep ## Run linter for the code
	golangci-lint run

cleanup-lint: dep ## Run all the linters and clean up
	golangci-lint run --fix

build-dev: dep ## Build a beta version
	go build -mod=vendor -race -o redesign_api .

build-docker: ## Build docker env
	docker-compose build

run-dev: build-dev ## Run the application locally
	./redesign_api api

test: ## Run test
	go test -race ./...

test-coverage: ## Run test coverage
	go test -v ./... -covermode=count -coverpkg=./... -coverprofile=coverage.out
	go tool cover -html coverage.out -o coverage.html

migrate: ## Apply outstanding migrations
	migrate

new-migration: ## New migration (make name=add-some-table new-migration)
	touch $(MIGRATION_FILE)
	echo "-- +migrate Up\n\n-- +migrate Down" >> $(MIGRATION_FILE)

start-env: ## Start the local env
	docker-compose up -d db

start-app: ## Start the application in docker container
	docker-compose up -d api

start-all: ## Start the environment services and the application in docker container
	docker-compose up -d

stop-env: ## Stop the local env
	docker-compose down

cleanup-build: ## Cleanup executable
	rm -f redesign_api

cleanup: cleanup-build ## Cleanup all files

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

connect-db: ## Connect to redesign local db
	psql postgresql://db:123@localhost:5442/redesign

.DEFAULT_GOAL := help
