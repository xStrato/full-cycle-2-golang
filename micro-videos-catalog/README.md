# micro-videos-catalog
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

### Generation mocks:
- for repositories:
```
mockgen -destination=mocks/repository.go -package=mock_interfaces github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces CategoryRepository
```
- for repositories(alt command):
```
mockgen -destination=mocks/repository.go -package=mock_interfaces github.com/xStrato/full-cycle-2-golang/micro-videos-catalog/domain/interfaces CategoryRepository,Repository,UnitOfWork
```
- for handlers:
```

```
