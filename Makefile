PHONY: run-server
run-server:
	go run ./api/main.go

PHONY: build
build:
	go build -o quiz-cli cmd/root.go


PHONY: quiz
quiz:
	go run main.go exam

PHONY: help
help:
	go run main.go -h

PHONY: test
test:
	go test -race -cover  ./... 

	