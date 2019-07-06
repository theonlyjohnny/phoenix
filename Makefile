BINARY_NAME=phoenix

all: test build 

build: 
	go build -o $(BINARY_NAME) -v ./cmd/$(BINARY_NAME)
test: 
	go test -race -v ./...
clean: 
	go clean ./...
	rm -f $(BINARY_NAME)
run:
	go run ./cmd/$(BINARY_NAME)
