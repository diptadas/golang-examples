### Run test

```
$ go test -v golang-examples/go_test -run ^TestSum$

=== RUN   TestSum
--- FAIL: TestSum (0.00s)
        math_test.go:22: Expected 5, got 6
        math_test.go:22: Expected 2, got 3
FAIL
exit status 1
FAIL    golang-examples/go_test 0.003s
```