DOCKER_GO=docker run -ti -v $PWD:/go/src/github.com/jlevesy/readstack:rw golang:alpine

all: build

build: create_dist
	@go build -o dist/server cmd/server/main.go

test:
	${DOCKER_GO} go test ./...

run:
	docker-compose up

create_dist:
	@mkdir -p dist

clean:
	@rm -rf dist
