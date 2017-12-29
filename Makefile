all: build

build: create_dist
	@go build -o dist/server cmd/server/main.go

create_dist:
	@mkdir -p dist

clean:
	@rm -rf dist
