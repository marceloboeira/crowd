DOCKER_COMPOSE  ?= `which docker-compose`
GO ?= `which go`
CROWD ?= `pwd`/bin/crowd

.PHONY: run
run:
	$(CROWD)

.PHONY: build
build:
	$(GO) build -o $(CROWD) crowd.go

.PHONY: compose
compose:
	$(DOCKER_COMPOSE) up

.PHONY: decompose
decompose:
	$(DOCKER_COMPOSE) down
