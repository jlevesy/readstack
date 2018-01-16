all: create_dist build web

.PHONY: build
build: vendor
	@go build -o dist/server cmd/server/main.go

.PHONY: test
test: unit_test integration_test

.PHONY: integration_test
integration_test:
	@echo "Running integration tests..."
	# @go test -v -race -timeout=10m -run=$(T) ./integration
	@echo "TODO..."

.PHONY: unit_test
unit_test:
	@echo "Running unit tests..."
	@go test -race -cover -timeout=5s -run=$(T) `go list ./... | grep -v integration`

.PHONY: clean_web
clean_web:
	@rm -rf dist/web

.PHONY: web
web: clean_web
	@echo "Building web appication"
	@mkdir -p dist/web
	@cp -r web/* dist/web
	@echo "Done !"

.PHONY: run_dev
run_dev:
	@docker-compose up

.PHONY: migate_up
migrate_up:
	@migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" up $(N)

.PHONY: migate_down
migrate_down:
	@migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" down $(N)

.PHONY: new_migration
new_migration:
	@migrate create -dir migration -ext sql $(NAME)

.PHONY: create_dist
create_dist:
	@mkdir -p dist

.PHONY: vendor
vendor:
	@dep ensure

.PHONY: clean
clean:
	@rm -rf dist
