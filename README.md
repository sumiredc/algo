# Algorithm in Golang


## Test 

### Coverage
```sh
# Export text
go test -cover ./... -coverprofile=coverage.out

# Export html
go tool cover -html=coverage.out -o coverage.html
```
