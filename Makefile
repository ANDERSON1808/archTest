test:
	go test ./...  -cover -coverprofile=coverage.out

check-linter:
	 golangci-lint run --path-prefix=./ -v --skip-dirs constant --config=./golangci-lint.yaml --timeout=5m