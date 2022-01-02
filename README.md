# full-cycle-2-golang
---
### Adopted development approaches:
 - DDD
 - CQRS
---
### Running Code Coverage:
```
go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o cover.html
```