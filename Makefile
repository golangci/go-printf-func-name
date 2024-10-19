.PHONY: lint test build vet

default: lint test vet

test:
	go test -v -cover ./...

lint:
	golangci-lint run

build:
	go build ./cmd/go-printf-func-name/

vet: build
	go vet -vettool=./go-printf-func-name ./...
