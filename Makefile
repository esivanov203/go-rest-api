.PHONY: build
build:
	go build -o ./apiserver ./cmd/apiserver

.DEFAULT_GOAL := build