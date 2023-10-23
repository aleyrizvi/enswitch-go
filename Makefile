lint:
	golangci-lint run --max-same-issues=0 --max-issues-per-linter=0

gofumpt:
	gofumpt -w ./...

coverage:
	go clean -testcache && go test -cover ./...

test:
	go clean -testcache && go test -v -cover