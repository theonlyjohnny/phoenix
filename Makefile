BINARY_NAME=phoenix
BINARY_LOCATION=phoenix-cli

AGENT_BINARY_NAME=phoenix-agent
AGENT_BINARY_LOCATOIN=phoenix-agent

all: test build 

build: 
	go build -o $(BINARY_NAME) -v ./cmd/$(BINARY_LOCATION)
	go build -o $(AGENT_BINARY_NAME) -v ./cmd/$(AGENT_BINARY_LOCATOIN)
test: 
	go test -race -v ./...
clean: 
	go clean ./...
	rm -f $(BINARY_NAME)
	rm -f $(AGENT_BINARY_LOCATOIN)
run:
	go run ./cmd/$(BINARY_LOCATION)
run-client:
	go run ./cmd/$(AGENT_BINARY_LOCATOIN)
