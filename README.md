# Server

## Quick start

``` shell
./start.sh
```

Then, visit the address http://localhost:9090/ping

> You can customize your configuration by edit `config/app.local.yml`

The swagger document located  

http://host/api/v1/swagger/index.html  
http://host/api/v1/swagger/doc.json

## Manual start

``` shell
# customize configurations
cp config/app.yml config/app.local.yml

# start postgres and redis service
docker-compose up -d

# bootstrap
go run main.go
```

## Test

``` shell
GIN_MODE=test go test
```
