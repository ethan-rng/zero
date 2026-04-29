.PHONY: build-web build test

build-web:
	cd internals/web && npm install && npm run build

build: build-web
	go build -o out/zero main.go

test:
	go test ./...
