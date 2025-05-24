.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY: fmt

lint: fmt
	staticcheck
.PHONY: lint

vet: fmt
	go vet ./...
.PHONY: vet

build: vet
	go mod tidy
.PHONY: build