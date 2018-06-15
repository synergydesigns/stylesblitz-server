build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphql lambda/graphql/main.go
	chmod +x lambda/bin/graphql
	env GOOS=linux go build -ldflags="-s -w -v" -o lambda/bin/graphqli lambda/graphqli/main.go
	chmod +x lambda/bin/graphqli
