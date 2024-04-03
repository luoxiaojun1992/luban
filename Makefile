PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: all
all: tidy vendor fmt lint test ui-install ui-fmt ui-lint

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
	bin/golangci-lint run

.PHONY: revive
revive: bin/revive
	bin/revive -config revive.toml -formatter friendly -exclude ./vendor/... ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: vet golangci-lint revive

.PHONY: test
test:
	go list ./... | xargs go test -cover -coverprofile=coverage.out

.PHONY: build
build:
	go build -o bin/luban cmd/luban/main.go

.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o bin/luban.wasm cmd/wasm/main.go

.PHONY: ui-install
ui-install:
	cd portal/luban-ui && npm install && cd ../../

.PHONY: ui-fmt
ui-fmt:
	cd portal/luban-ui && npm run format && cd ../../

.PHONY: ui-lint
ui-lint:
	cd portal/luban-ui && npm run lint && cd ../../

.PHONY: ui-run-dev
ui-run-dev:
	cd portal/luban-ui && npm run dev

.PHONY: build-ui
build-ui:
	cd portal/luban-ui && npm run build && cd ../../
