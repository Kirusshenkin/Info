# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get
BINARY_NAME=main

# Make commands
.PHONY: all test clean tidy build run

# Default command
all: build

# Build the project
build:
	$(GOBUILD) -o $(BINARY_NAME) ./cmd

# Run the project
run:
	./$(BINARY_NAME)

# Clean up
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run tests
test:
	$(GOTEST) -v ./...

# Update dependencies
tidy:
	$(GOMOD) tidy

# Clean module cache
clean-modcache:
	$(GOCLEAN) -modcache

# Download modules
mod-download:
	$(GOMOD) download

# Get cryptocurrency prices
crypto:
	go run ./cmd/main.go crypto

# Additional commands if needed
docker-build:
	docker build -t your-image-name .

docker-run:
	docker run -p 8080:8080 your-image-name
