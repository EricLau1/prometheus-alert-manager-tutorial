# Golang With Prometheus Alert Manager


## Config Swagger

- REF:
  - https://santoshk.dev/posts/2022/how-to-integrate-swagger-ui-in-go-backend-gin-edition/
  - https://github.com/swaggo/swag/tree/master/example/celler/controller

### Dependencies:

```bash
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/files
go get -u github.com/swaggo/gin-swagger
```

### Generate Swagger

```bash
swag init
```

Output...

```bash
2022/08/17 21:34:06 Generate swagger docs....
2022/08/17 21:34:06 Generate general API Info, search dir:./
2022/08/17 21:34:08 Generating types.Todo
2022/08/17 21:34:08 Generating httpext.JsonError
2022/08/17 21:34:08 create docs.go at docs/docs.go
2022/08/17 21:34:08 create swagger.json at docs/swagger.json
2022/08/17 21:34:08 create swagger.yaml at docs/swagger.yaml
```

## Run Docker

```bash
cd docker

make up
```

