.PHONY: clean

default: build

prepare:
	go mod tidy

build: prepare
	go build -o mcp
run:
	./mcp
