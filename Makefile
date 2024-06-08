BIN_DIR := $(CURDIR)/bin

STATICCHECK := $(BIN_DIR)/staticcheck
TESTIFYILINT := $(BIN_DIR)/testifylint


$(STATICCHECK):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install honnef.co/go/tools/cmd/staticcheck@latest

$(TESTIFYILINT):
	mkdir -p $(BIN_DIR)
	GOBIN=$(BIN_DIR) go install github.com/Antonboom/testifylint@latest

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: check
check: $(STATICCHECK) $(TESTIFYILINT)
	$(STATICCHECK) ./...
	go vet ./...
	$(TESTIFYILINT) ./...

.PHONY: test
test:
	go test -v ./...
