.PHONY: run build

test-run:
	go run cmd/filepatrol/main.go --path tmp --cmd 'echo "hello world"'

test-run-2:
	go run cmd/filepatrol/main.go --cmd "jq '.items[-1]' ./logs/errors.json ; jq '. | length' ./logs/errors.json" --path ./logs

test-run-http:
	go run cmd/filepatrol/main.go --path tmp --exec http

build:
	mkdir -p build
	go build -o build/filepatrol
