# Code Challenge: Mutants


## Instructions on how to run the project

The next commands bellow will run the project.

#### Local

```bash
go mod tidy
go run ./cmd/api/main.go
```

#### Dockerized

```bash
make dev
```

## Others

### Code coverage

```bash
go test ./... -cover
```

### Lint

```bash
golint ./...
```

## Extra documentation

### Swagger

The application includes [swagger](./api/docs/swagger.yaml) definition.

### Postman

The application includes potman [Collection](./api/postman/meli_test.postman_collection.json)
