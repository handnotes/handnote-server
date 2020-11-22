# Server

## Getting Start

``` bash
cp .env.example .env
./start.sh
```

Then, visit the address http://localhost:9090/hello

> The swagger document located http://localhost:9090/api/v1/swagger/index.html

## Initial Postgres

``` bash
docker-compose up -d postgres
```

## Initial Redis

``` bash
docker-compose up -d redis
```

## Setup SMTP service

Modify the SMTP server and password in `.env` file

## Test
``` bash
# set test env
$ GIN_MODE=test go test
```
