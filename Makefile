.PHONY: run build

test-run:
	go run cmd/filepatrol/main.go --path tmp --cmd echo "hello"

test-run-http:
	go run cmd/filepatrol/main.go --path tmp --exec filepatrol.http

build:
	mkdir -p build
	go build -o build/filepatrol
