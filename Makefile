.PHONY: run build

run:
	go run ./cmd/filepatrol/*.go

build:
	mkdir -p build
	go build -o build/go-file-watcher
