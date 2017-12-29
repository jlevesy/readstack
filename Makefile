all: build

build: create_dist
	@go build -o dist/server cmd/server/main.go

test:
	@go test ./...

run:
	@docker-compose up

migrate_up:
	@migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" up $(N)

migrate_down:
	@migrate -path ./migration -database "postgres://readstack:notsecret@localhost:5432/readstack?sslmode=disable" down $(N)

new_migration:
	@migrate create -dir migration -ext sql $(NAME)

create_dist:
	@mkdir -p dist

clean:
	@rm -rf dist
