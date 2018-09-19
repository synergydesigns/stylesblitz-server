TEST_PACKAGES := $(shell go list ./shared/... | grep -v vendor | grep -v fakes)


build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphql lambda/graphql/main.go
	chmod +x lambda/bin/graphql
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphqli lambda/graphqli/main.go
	chmod +x lambda/bin/graphqli

test:
	$(shell set -o allexport source .env set +o allexport)
	echo TEST_PACKAGES
	@printenv
	@go test -v -cover $(TEST_PACKAGES)

migrate:
	@go run shared/migrations/main.go