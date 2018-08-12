all: create_dist build

.PHONY: static_build
static_build: vendor
	@./toolbox @CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags -static' -o dist/server backend/cmd/server/main.go

.PHONY: build
build: vendor
	@./toolbox go build -o dist/server cmd/server/main.go

.PHONY: test
test: unit_test integration_test

.PHONY: vendor
vendor:
	@./toolbox dep ensure

.PHONY: integration_test
integration_test:
	@echo "Running integration tests..."
	# @go test -v -race -timeout=10m -run=$(T) ./integration
	@echo "TODO..."

.PHONY: unit_test
unit_test:
	@echo "Running unit tests..."
	@./toolbox go test -race -cover -timeout=5s -run=$(T) `go list ./... | grep -v test`

.PHONY: clean_web
clean_web:
	@./toolbox rm -rf dist/web

.PHONY: run_dev
run_dev:
	@docker-compose up

.PHONY: migate_up
migrate_up:
	@./toolbox migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" up $(N)

.PHONY: migate_down
migrate_down:
	@./toolbox migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" down $(N)

.PHONY: new_migration
new_migration:
	@./toolbox migrate create -dir migration -ext sql $(NAME)

.PHONY: create_dist
create_dist:
	@./toolbox mkdir -p dist

.PHONY: clean
clean:
	@./toolbox rm -rf dist

.PHONY: toolbox
toolbox:
	docker build --build-arg=UID=$(id -u) --build-arg=GID=$(id -g) -t readstack-toolbox -f Dockerfile.toolbox .
