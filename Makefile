DOCKER_COMPOSE  ?= `which docker-compose`
GO ?= `which go`
DEP ?= `which dep`
CROWD_ENTRYPOINT ?= `pwd`/src/crowd.go
CROWD ?= `pwd`/bin/crowd

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
