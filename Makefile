all: test

.PHONY: test
test:
	go get -t -v ./...
	go test -v -coverprofile=coverage.out -covermode=count
	go tool cover -html=coverage.out
