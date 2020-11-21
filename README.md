# Server

## Getting Start

``` bash
source ./sendgrid.env
go run handnote.go
```

Then, visit the address http://localhost:9090/hello

## Initial Postgres

``` bash
docker-compose up -d postgres
```

## Initial Redis

``` bash
$ brew install redis
$ brew services start redis
$ redis-cli

keys *
```

## OpenAPI

``` bash
# go-swagger generate api doc and serve it.
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ swagger generate spec -o ./swagger.yml
$ swagger serve ./swagger.yml
```

## Test
``` bash
# set test env
$ GIN_MODE=test go test
```
