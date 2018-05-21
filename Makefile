all: test

.PHONY: test
test:
	go test -v -coverprofile=coverage.out -covermode=count
	go tool cover -html=coverage.out
