VERSION ?= 0.0.1
COMMIT ?= $$(git describe --always)
BINARY_NAME ?= gstar

default: build

clean:
	rm $(GOPATH)/bin/$(BINARY_NAME)
	rm bin/$(BINARY_NAME)

deps:
	go get -d -t -v .

build: deps
	go build -ldflags "-X main.GitCommit=$(COMMIT)" -o bin/$(BINARY_NAME)

install: deps
	go install -ldflags "-X main.GitCommit=$(COMMIT)"

test-all: vet test

test:
	go test -v -parallel=4 .

test-race:
	@go test -race .

vet:
	go vet *.go

lint:
	@go get github.com/golang/lint/golint
	golint ./...

# cover shows test coverages
cover:
	@go get golang.org/x/tools/cmd/cover
	godep go test -coverprofile=cover.out
	go tool cover -html cover.out
	rm cover.out