# balance-api-go

## Development

Build the project

```
make build
```

Run the project

```
make run
```

Generate mocks with `mockgen` 
``go install github.com/golang/mock/mockgen@v1.6.0``

```
make generate-mocks
```

### Testing

Run unit tests

```
make unit-test
```

Run lint

```
make lint
```

Fix lint

```
make lint-fix
```

Run all tests

```
make all-tests
```

### Accessing Service

Service will be automatically forwarded to `localhost:3001`
