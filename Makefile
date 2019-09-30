export CC_TEST_REPORTER_ID = 98874d12dfe651a7dfa0c452370ba52f7c92bd55a74ebc2525d3fcef90fd22f6
COVERAGEOUTFILE=c.out

all: test race

.PHONY: dep
dep:
	go mod vendor

.PHONY: test-vendor
test-vendor:
	go test -mod=vendor -coverprofile=coverage.out -covermode=count

.PHONY: test
test: dep
	go test -coverprofile=coverage.out -covermode=count

.PHONY: race
race: dep
	go test -race

.PHONY: test-report
test-report: test
	go tool cover -html=coverage.out

.PHONY: travis
travis:
	# install deps pre-1.13
	go get -u github.com/google/go-cmp/cmp
	go test -coverprofile $(COVERAGEOUTFILE) ./...

.PHONY: cclimate-linux
cclimate-linux:
	curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-linux-amd64 > ./cc-test-reporter
	# curl -L https://codeclimate.com/downloads/test-reporter/test-reporter-latest-darwin-amd64 > ./cc-test-reporter
	chmod +x ./cc-test-reporter
	./cc-test-reporter before-build
	# install deps pre-1.13
	go get -u github.com/google/go-cmp/cmp
	go test -coverprofile $(COVERAGEOUTFILE) ./...
	./cc-test-reporter after-build --exit-code $(TRAVIS_TEST_RESULT)
