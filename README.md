# full-cycle-2-golang

### Running Code Coverage:
```
go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```