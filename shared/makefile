TEST_PACKAGES := $(shell go list ./... | grep -v vendor | grep -v fakes)

test:
	echo TEST_PACKAGES
	@go test -v -cover $(TEST_PACKAGES)