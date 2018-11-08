TEST_PACKAGES := $(shell go list ./shared/... | grep -v vendor | grep -v fakes)


build:
	dep ensure
	make recompile

test:
	@go test -v -cover $(TEST_PACKAGES)

migrate:
 $(shell cd shared/migrations && go run main.go)

recompile:
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphql lambda/graphql/main.go
	chmod +x lambda/bin/graphql
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphqli lambda/graphqli/main.go
	chmod +x lambda/bin/graphqli

dev-recompile:
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphql lambda/graphql/main.go
	chmod +x lambda/bin/graphql

supervise:
	supervisor --no-restart-on exit -e go -i bin --exec make -- dev-recompile

start-local:
	sam local start-api --profile serverless-pm --env-vars env.json