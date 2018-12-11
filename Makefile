DOCKER_COMPOSE=`which docker-compose`

.PHONY: run
run:
	echo RUN

.PHONY: compose
compose:
	$(DOCKER_COMPOSE) up

.PHONY: decompose
decompose:
	$(DOCKER_COMPOSE) down
