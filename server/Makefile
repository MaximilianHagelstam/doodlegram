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

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
