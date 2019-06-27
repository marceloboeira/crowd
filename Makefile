DOCKER_COMPOSE  ?= `which docker-compose`
GO ?= `which go`
DEP ?= `which dep`
CROWD_ENTRYPOINT ?= `pwd`/src/crowd.go
CROWD ?= `pwd`/bin/crowd

.PHONY: help
help: ## Lists the available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Builds the app main binary
	@$(GO) build -o $(CROWD) $(CROWD_ENTRYPOINT)

.PHONY: install
install: ## Installs app dependencies
	@$(DEP) ensure

.PHONY: run
run: ## Runs the app
	@$(CROWD)

.PHONY: test
test: ## Runs tests
	@$(GO) test -v ./...

.PHONY: compose
compose: ## Runs docker compose dependencies
	@$(DOCKER_COMPOSE) up

.PHONY: decompose
decompose: ## Stops docker compose dependencies
	@$(DOCKER_COMPOSE) down
