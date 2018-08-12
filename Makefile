.PHONY: run_dev
run_dev:
	@docker-compose up

.PHONY: toolbox
toolbox:
	@docker build --build-arg=UID=$(shell id -u) --build-arg=GID=$(shell id -g) -t readstack-toolbox -f Dockerfile.toolbox .
