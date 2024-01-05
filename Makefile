.PHONY: run build

test-run:
	go run cmd/filepatrol/main.go tmp echo 'hello'

build:
	mkdir -p build
	go build -o build/filepatrol
