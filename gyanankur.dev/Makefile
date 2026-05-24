.PHONY: run build docker test

run:
	go run ./cmd/server

build:
	go build -o bin/server ./cmd/server

docker:
	docker build -t gyanankur-dev .

docker-run:
	docker run -p 8080:8080 gyanankur-dev

test:
	go test ./...
