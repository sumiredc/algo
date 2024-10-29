# Algorithm in Golang
![coverage](https://raw.githubusercontent.com/sumiredc/algo/badges/.badges/main/coverage.svg)

## Test 

### Coverage
- https://github.com/marketplace/actions/go-test-coverage

```sh
# Export text
go test -cover ./... -coverprofile=coverage.out

# Export html
go tool cover -html=coverage.out -o coverage.html
```

## Lint
- https://golangci-lint.run/

```sh
golangci-lint run ./...
```
