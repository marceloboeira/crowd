DOCKER_COMPOSE  ?= `which docker-compose`
GO ?= `which go`
DEP ?= `which dep`
CROWD_ENTRYPOINT ?= `pwd`/src/crowd.go
CROWD ?= `pwd`/bin/crowd

.PHONY: help
help: ## Lists the available commands. Add a comment with '##' to describe a command.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build:
	@$(GO) build -o $(CROWD) $(CROWD_ENTRYPOINT)

.PHONY: install
install:
	@$(DEP) ensure

.PHONY: run
run:
	@$(CROWD)

.PHONY: test
test:
	@$(GO) test -v ./...

.PHONY: compose
compose:
	@$(DOCKER_COMPOSE) up

.PHONY: decompose
decompose:
	@$(DOCKER_COMPOSE) down
