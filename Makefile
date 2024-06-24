.PHONY: test-run build test-run-http

build:
	mkdir -p build
	go build -o ./build/filepatrol cmd/filepatrol/main.go

test-run:
	go run ./cmd/filepatrol/main.go --path tmp --cmd 'echo "hello world"'

test-run-http:
	go run cmd/filepatrol/main.go --path tmp --exec http

list:
	GOPROXY=proxy.golang.org go list -m github.com/sifatulrabbi/filepatrol@v0.3.0-beta.2
