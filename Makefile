# Define variables
DOCKER_COMPOSE = docker compose
DOCKER = docker

# Containers
ENGINE_CONTAINER_NAME = $(shell docker ps | grep engine- | rev | cut -d' ' -f1 | rev)
REDIS_CONTAINER_NAME = $(shell docker ps | grep redis | rev | cut -d' ' -f1 | rev)


# Default target
all: setup

##@ General setup

.PHONY: setup
setup: ## Build and start the Docker containers
	cp .env.example .env
	$(DOCKER_COMPOSE) build
	$(DOCKER_COMPOSE) up -d


##@ Engine

.PHONY: reset-world
reset-world: ## Delete all world data from redis and restart engine, execute when the project is running
	$(DOCKER_COMPOSE) stop engine
	$(DOCKER) exec -it $(REDIS_CONTAINER_NAME) redis-cli FLUSHALL
	$(DOCKER_COMPOSE) start engine
	@echo "Game world was erased!"

.PHONY: attach-engine
attach-engine: ## Attach to the engine container
	$(DOCKER) attach $(ENGINE_CONTAINER_NAME)

.PHONY: gm-set
gm-set: ## Set player with ID as Game Master
	$(DOCKER_COMPOSE) stop engine
	$(DOCKER_COMPOSE) run --rm engine go run cmd/cli/main.go gm:set --goid=$(ID)
	$(DOCKER_COMPOSE) start engine

##@ Everyday usage

.PHONY: start
start: ## Start all Docker containers
	$(DOCKER_COMPOSE) up -d

start-debug: docker-compose.yml docker-compose.debug.yml ## Start all Docker containers in debug mode
	$(DOCKER_COMPOSE) -f docker-compose.yml -f docker-compose.debug.yml up

.PHONY: stop
stop: ## Stop all Docker containers
	$(DOCKER_COMPOSE) stop

.PHONY: restart
restart: ## Restart all Docker containers
	$(MAKE) stop
	$(MAKE) start

.PHONY: clean
clean: ## Stop and remove Docker containers
	$(DOCKER_COMPOSE) down
	$(DOCKER_COMPOSE) rm -f


##@ Other

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
