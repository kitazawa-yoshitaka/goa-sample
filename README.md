# goa-sample
[goa](https://github.com/goadesign/goa)のお試しリポジトリ


## Environment
Generate code.
```
$ goa gen goa-sample/design
```

Run server.
```
$ cd cmd/calc
$ go build
$ ./calc
```

Curl request
```
$ curl http://localhost:8080/add/1/2
3
```