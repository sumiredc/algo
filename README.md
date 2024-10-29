# Algorithm in Golang


## Test 

### Coverage
- https://github.com/marketplace/actions/go-test-coverage

```sh
# Export text
go test -cover ./... -coverprofile=coverage.out

# Export html
go tool cover -html=coverage.out -o coverage.html
```
