BINARY_NAME=phoenix
BINARY_LOCATION=phoenix-cli

all: test build 

build: 
	go build -o $(BINARY_NAME) -v ./cmd/$(BINARY_LOCATION)
test: 
	go test -race -v ./...
clean: 
	go clean ./...
	rm -f $(BINARY_NAME)
run:
	go run ./cmd/$(BINARY_LOCATION)
