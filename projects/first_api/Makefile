
.PHONY: all
all: _build

DOCKER_COMPOSE=docker-compose  -f deployments/docker-compose.yml --project-name first_api

_build:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE) up  blog_initialization
	$(DOCKER_COMPOSE) stop

install:
	$(DOCKER_COMPOSE) stop
	$(DOCKER_COMPOSE) up -d blog_service

uninstall:
	$(DOCKER_COMPOSE) stop

.PHONY: clean

clean:
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) rm -f
