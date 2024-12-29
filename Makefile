default: build test

build: protoc
	go build -o $GOPATH/bin/copper/ ./cmd/*

test:
	go clean --testcache
	go test ./test/...

protoc:
	protoc --proto_path="./proto" --go_out="./" --go-grpc_out="./" ./proto/*.proto