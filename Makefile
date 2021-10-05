api_cmd = surt-api
controller_cmd = surt-controller
runner_cmd = surt-task-runner

install:
	@echo "> Installing all required tools..."
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install github.com/goreleaser/goreleaser@latest
	@go install github.com/nektos/act@latest

mods:
	@echo "> Downloading all required go modules..."
	@go mod download

verify: 
	@go mod verify

lint:
	staticcheck ./...

test:
	go test ./pkg/...

coverage:
	@echo "> Running coverage tests..."
	go test -coverpkg=./pkg/... -coverprofile=coverage.out ./pkg/...

build:
	@echo "> Bulding all binaries..."
	goreleaser release --snapshot --skip-publish --rm-dist

clamav-start:
	@echo "> Starting local clamav using docker..."
	docker run -d -p 3310:3310 --rm --name clamav ghcr.io/surt-io/container-clamav:latest-initdb

clamav-stop:
	@echo "> Stopping local clamav..."
	docker stop clamav

clamav-remove:
	@echo "> Removing local clamav container..."
	docker rm -v -f clamav

all: install mods verify lint test build

act-pr:
	@echo "> Running Github Actions locally for PR event..."
	act -s GITHUB_TOKEN pull_request
