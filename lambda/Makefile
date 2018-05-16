build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/graphql graphql/main.go
	chmod +x bin/graphql
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/graphqli graphqli/main.go
	chmod +x bin/graphqli
