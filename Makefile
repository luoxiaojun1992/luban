PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: tidy vendor fmt lint test

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

bin/goimports:
	mkdir -p temp && cd temp && go mod init temp && GOBIN=$(PROJECT_DIR)/bin go install golang.org/x/tools/cmd/goimports@latest && cd .. && rm -rf temp

.PHONY: goimports
goimports: bin/goimports
	bin/goimports -w .

.PHONY: fmt
fmt: goimports
	go fmt ./...

bin/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.57.2

bin/revive:
	mkdir -p temp && cd temp && go mod init temp && GOBIN=$(PROJECT_DIR)/bin go install github.com/mgechev/revive@v1.3.7 && cd .. && rm -rf temp

.PHONY: golangci-lint
golangci-lint: bin/golangci-lint
	bin/golangci-lint run --skip-dirs=data/test_stubs

.PHONY: revive
revive: bin/revive
	bin/revive -config revive.toml -formatter friendly -exclude ./vendor/... -exclude ./data/test_stubs/... ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: vet golangci-lint revive

.PHONY: test
test:
	go list ./... | grep -v "data/test_stubs" | xargs go test -cover -coverprofile=coverage.out

.PHONY: build
build:
	go build -o bin/luban cmd/luban/main.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o bin/luban.wasm cmd/wasm/main.go
