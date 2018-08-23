.PHONY: run_dev
run_dev:
	@docker-compose up

.PHONY: generate_go
generate_go:
	@mkdir -p server/api
	@protoc -I interface/ interface/readstack.proto --go_out=plugins=grpc:server/api

.PHONY: generate_php
generate_php:
	@mkdir -p php/src/Readstack
	@protoc --php_out=php/src/Readstack --grpc_out=php/src/Readstack --plugin=protoc-gen-grpc=$(HOME)/grpc/bins/opt/grpc_php_plugin ./interface/readstack.proto

.PHONY: toolbox
toolbox:
	@docker build --build-arg=UID=$(shell id -u) --build-arg=GID=$(shell id -g) -t readstack-toolbox -f Dockerfile.toolbox .
