# Server

## Getting Start

```bash
$ git clone -b server https://e.coding.net/handnote/handnote.git
$ cd handnote
$ source ./sendgrid.env
$ go run main.go
# Then, visit the address http:/localhost:9090/hello.
```

## OpenAPI

```bash
# go-swagger generate api doc and serve it.
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
$ swagger generate spec -o ./swagger.yml
$ swagger serve ./swagger.yml
```
