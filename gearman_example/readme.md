### Gearman Example
[Gearman](http://gearman.org/) client and worker implementation in [golang](http://golang.org) using [appscode/g2](https://github.com/appscode/g2).

### Install requirements

```
$ go get -v github.com/appscode/g2
$ go get -v github.com/mikespook/golib/signal
$ go get -v golang.org/x/tools/cmd/goimports

$ ./go/src/github.com/appscode/g2/hack/builddeps.sh
$ ./go/src/github.com/appscode/g2/hack/make.py
```

### Run server

```
$ gearmand --v=10 --addr="0.0.0.0:4730"
or
$ go run go/src/github.com/appscode/g2/cmd/gearmand/main.go --v=10 --addr="0.0.0.0:4730"
```

### Run worker

```
$ go run go/src/gearman-example/worker/worker.go
```

### Run client

```
$ go run go/src/gearman-example/client/client.go
```
