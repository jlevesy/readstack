all: create_dist build

.PHONY: static_build
static_build: vendor
	@CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags -static' -o dist/server backend/cmd/server/main.go

.PHONY: build
build: vendor
	@go build -o dist/server cmd/server/main.go

.PHONY: test
test: unit_test integration_test

.PHONY: vendor
vendor:
	@dep ensure

.PHONY: integration_test
integration_test:
	@echo "Running integration tests..."
	# @go test -v -race -timeout=10m -run=$(T) ./integration
	@echo "TODO..."

.PHONY: unit_test
unit_test:
	@echo "Running unit tests..."
	@go test -race -cover -timeout=5s -run=$(T) `go list ./... | grep -v test`

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
