.PHONY: lint test

default: lint test

test:
	go test -v -cover ./...

lint:
	golangci-lint run
