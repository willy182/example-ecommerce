# Example-Ecommerce

## Prepare service
```
$ docker compose up
```

```
$ go mod tidy
```

### If using SQL database, run this commands for migration
UP migration:
```
$ make migration
```

Rollback migration:
```
$ make migration down
```

## Run unit test & calculate code coverage

Make sure generate mock using [mockery](https://github.com/vektra/mockery)
```
$ make mocks
```

Run test:
```
$ make test
```

## Build and run service
```
$ make run
```

## Dashboard Monitor
http://localhost:8080

## Postman Collection
```
in folder docs
```