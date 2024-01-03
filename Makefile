.PHONY: run build

run:
	go run ./*.go

build:
	mkdir -p build
	go build -o build/go-file-watcher
