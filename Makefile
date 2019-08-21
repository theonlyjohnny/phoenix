CLI_NAME=phoenix
CLI_LOCATION=phoenix-cli

AGENT_NAME=phoenix-agent
AGENT_LOCATION=phoenix-agent

DIST=./dist

TEST_OUTPUT=/tmp/phoenix/test_output

test:
	mkdir -p $(TEST_OUTPUT)
	TEST_OUTPUT=$(TEST_OUTPUT) go test -race ./... -coverprofile=$(TEST_OUTPUT)/phoenix_coverage.out -covermode=atomic

coverage: test
	go tool cover -html=$(TEST_OUTPUT)/phoenix_coverage.out

clean:
	go clean ./...
	rm -f $(CLI_NAME)
	rm -f $(AGENT_LOCATION)
	rm -rf $(DIST)
	mkdir $(DIST)

run:
	go run ./cmd/$(CLI_LOCATION)
run-client:
	go run ./cmd/$(AGENT_LOCATION)

run-all:
	go run ./cmd/$(AGENT_LOCATION) &
	go run ./cmd/$(CLI_LOCATION)


build: clean
	go build -o $(DIST)/$(CLI_NAME) -v ./cmd/$(CLI_LOCATION)
	go build -o $(DIST)/$(AGENT_NAME) -v ./cmd/$(AGENT_LOCATION)

all: test build
