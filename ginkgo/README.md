### Ginkgo Example

[Ginkgo](https://github.com/onsi/ginkgo) is a [BDD](https://en.wikipedia.org/wiki/Behavior-driven_development) style testing framework for Golang.

```
// Install
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega

// Bootstrap suit
$ cd go/src/golang-examples/ginkgo_example
$ ginkgo bootstrap

// Generate specs
$ ginkgo generate book

// Run tests after adding assertions to specs
$ ginkgo -v
```