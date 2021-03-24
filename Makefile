
GOPATH=$(shell go env GOPATH)
GOSTATIC=--trimpath --tags "netgo sqlite_omit_load_extension sqlite_foreign_keys" -ldflags '-linkmode external -extldflags "-static" -s -w'


test:
	go test --tags "sqlite_foreign_keys" -v ./...

build:
	go build $(GOSTATIC) -o bin/agent github.com/ubogdan/gopanel/cmd/agent
	go build $(GOSTATIC) -o bin/server github.com/ubogdan/gopanel/cmd/server

fmt:
	go fmt ./...