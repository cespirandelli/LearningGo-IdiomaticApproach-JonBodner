.DEFAULT_GOAL := build

fmt:
		go fmt ./...
.PHONY:fmt

lint: fmt
		golint ./...
.PHONY:lint

vet:
		go vet  ./...
.PHONY:vet

build:
		go build hello.go
.PHONY:build