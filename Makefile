export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: bins, version.o, test

BIN_DIR = bins

default: bins

lint: linter
	$(BIN_DIR)/golangci-lint run ./...

linter:
	@sh ./scripts/install_golangci.sh $(BIN_DIR)

lint-fix: linter install-tools
	goimports -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)
	$(BIN_DIR)/golangci-lint run --fix ./...

install-tools:
	@echo "Installing Tools from tools.go"
	@go list -f '{{join .Imports "\n"}}' ./tools.go | xargs  go install