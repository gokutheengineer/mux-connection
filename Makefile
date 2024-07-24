SHELL=/bin/bash

.PHONY: deps
deps:
	go mod download

.PHONY: build
build:
	go build -o bin/ ./...

.PHONY: run
run:
	go run *.go