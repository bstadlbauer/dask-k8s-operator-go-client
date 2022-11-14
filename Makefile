.PHONY: setup
setup:
	go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0
	go install github.com/go-critic/go-critic/cmd/gocritic@v0.6.5
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.1
	go install golang.org/x/tools/cmd/goimports@v0.2.0
	pre-commit install
