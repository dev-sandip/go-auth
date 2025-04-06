APP_NAME = app
CMD_DIR = cmd/app
BIN_DIR = bin

build:
	@echo "🔧 Building $(APP_NAME)..."
	go build -o $(BIN_DIR)/$(APP_NAME) ./$(CMD_DIR)

run: build
	@echo "🚀 Running $(APP_NAME)..."
	./$(BIN_DIR)/$(APP_NAME)

generate:
	templ generate

dev:
	air

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BIN_DIR)
