all: test

.PHONY: test
test:
	dep ensure
	go test -v -coverprofile=coverage.out -covermode=count
	go tool cover -html=coverage.out
