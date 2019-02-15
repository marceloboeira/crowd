DOCKER_COMPOSE  ?= `which docker-compose`
GO ?= `which go`
CROWD ?= `pwd`/bin/crowd
VEGETA ?= `which vegeta`
VEGETA_FOLDER ?= `pwd`/performance
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

.PHONY: performance
performance:
	cat $(VEGETA_FOLDER)/Vegetafile | $(VEGETA) attack -duration=5m  -rate=050 | $(VEGETA) report -reporter=plot -output=$(VEGETA_FOLDER)/report_050.html
	cat $(VEGETA_FOLDER)/Vegetafile | $(VEGETA) attack -duration=5m  -rate=100 | $(VEGETA) report -reporter=plot -output=$(VEGETA_FOLDER)/report_100.html
	cat $(VEGETA_FOLDER)/Vegetafile | $(VEGETA) attack -duration=5m  -rate=500 | $(VEGETA) report -reporter=plot -output=$(VEGETA_FOLDER)/report_500.html
