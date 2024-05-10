# Build the application
all: build

build:	
	@go build -o main cmd/main.go

# Run the application
run:
	@go run cmd/main.go

# Create DB container
db-up:
	@docker compose up

# Shutdown DB container
db-down:
	@docker compose down

# Test the application
test:
	@go test ./tests -v

# Clean the binary
clean:
	@rm -f main

# Live reload
watch:
	@air

# Lint
lint:
	@golangci-lint run

.PHONY: all build run test clean lint
