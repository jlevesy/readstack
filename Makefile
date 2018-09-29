all: create_dist build

.PHONY: run_dev
run_dev:
	@docker-compose up

.PHONY: toolbox
toolbox: cachedirs
	@docker build --build-arg=UID=$(shell id -u) --build-arg=GID=$(shell id -g) -t readstack-toolbox -f Dockerfile.toolbox .

.PHONY: cachedirs
cachedirs:
	@mkdir -p .gocache/mod
	@mkdir -p .gocache/build

.PHONY: generate_go
generate_go:
	@protoc -I api/ api/readstack.proto --go_out=plugins=grpc:api

.PHONY: install
install:
	go install cmd/readstackctl/main.go

.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags -static' -o dist/server cmd/server/main.go

.PHONY: test
test: unit_test integration_test

.PHONY: integration_test
integration_test:
	@echo "Running integration tests..."
	@go test -v -race -timeout=10m -run=$(T) ./integration

.PHONY: unit_test
unit_test:
	@echo "Running unit tests..."
	@go test -race -cover -timeout=5s -run=$(T) `go list ./... | grep -v integration`

.PHONY: migate_up
migrate_up:
	@sql-migrate up

.PHONY: migate_down
migrate_down:
	@sql-migrate down

.PHONY: new_migration
new_migration:
	@sql-migrate new

.PHONY: create_dist
create_dist:
	@mkdir -p dist

.PHONY: clean
clean:
	@rm -rf dist
