> Note: This operator has been moved over to [dask/dask-kubernetes](https://github.com/dask/dask-kubernetes). Please use this one instead! 

# dask-k8s-operator-go-client
A go client for the dask kubernetes operator



## Contributing
### Prerequisites
- [go](https://go.dev/) >= 1.18. Might work with an older version, but not tested.
- [pre-commit](https://pre-commit.com/index.html). The `go` hooks will require the following tools:
    - [gocyclo](https://github.com/fzipp/gocyclo)
    - [go-critic](https://github.com/go-critic/go-critic)
    - [golanclii-lint](https://github.com/golangci/golangci-lint)
    - [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)

To set all of these up for you, you can run: 

```shell
make setup
```

