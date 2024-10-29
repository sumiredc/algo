# Algorithm in Golang


## Test 

### Coverage
```sh
# Export text
go test -cover ./... -coverprofile=cover.out

# Export html
go tool cover -html=cover.out -o cover.html 
```
