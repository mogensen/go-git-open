.PHONY: help update-protobuf docker-build test-coverage build

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test-coverage: ## Generate test coverage report
	mkdir -p tmp
	go test ./... --coverprofile tmp/outfile
	go tool cover -html=tmp/outfile

report-card: ## Generate static analysis report
	goreportcard-cli -v

build: ## Builds a static linked binary
	go build -o git-open ./cmd/git-goopen

install: ## Installs the static linked binary
	go install ./cmd/git-goopen
