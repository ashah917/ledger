BINARY_NAME=ledger
BINARY_DIR=bin

.PHONY: build clean run

build:
	mkdir -p $(BINARY_DIR)
	go build -o $(BINARY_DIR)/$(BINARY_NAME) ledger_provider.go block_provider.go name_server.go translater.go

clean:
	rm -rf $(BINARY_DIR)

run: build
	./$(BINARY_DIR)/$(BINARY_NAME)

