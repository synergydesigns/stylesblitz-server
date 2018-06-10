HANDLERS = $(shell find . -name "main.go" | grep -v vendor | grep -v bin; ) 

build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/graphql graphql/main.go
	chmod +x bin/graphql
	env GOOS=linux go build -ldflags="-s -w -v" -o bin/graphqli graphqli/main.go
	chmod +x bin/graphqli

# test:
# 	$(shell for i in HANDLERS; \
# 		do \
# 			dep ensure; \
# 			env GOOS=linux go build -ldflags="-s -w -v" -o bin/graphql graphql/main.go; \
# 			chmod +x bin/graphql; \
# 		done)
