export PATH := $(GOPATH)/bin:$(PATH)

.PHONY: bins, version.o, test

BIN_DIR = bins

default: bins

lint: linter
	$(BIN_DIR)/golangci-lint run ./...

linter:
	@sh ./scripts/install_golangci.sh $(BIN_DIR)

lint-fix: linter
	goimports -w $$(go list -f {{.Dir}} ./... | grep -v /vendor/)
	$(BIN_DIR)/golangci-lint run --fix ./...
