# go-deque

[![Go Reference](https://pkg.go.dev/badge/github.com/jayp0521/go-deque.svg)](https://pkg.go.dev/github.com/jayp0521/go-deque)
[![Go Report Card](https://goreportcard.com/badge/github.com/jayp0521/go-deque)](https://goreportcard.com/report/github.com/jayp0521/go-deque)

# Benchmark

```
BenchmarkAppendFirst-16                 20438316                65.57 ns/op           24 B/op          1 allocs/op
BenchmarkAppendFirst_List-16            100000000               30.27 ns/op           45 B/op          0 allocs/op

BenchmarkAppendPopFirst-16              17881148                78.14 ns/op           24 B/op          1 allocs/op
BenchmarkAppendPopFirst_List-16         100000000               32.96 ns/op           45 B/op          0 allocs/op

BenchmarkInsertAt-16                    20175003                66.46 ns/op           24 B/op          1 allocs/op
BenchmarkInsertAt_List-16               100000000               34.19 ns/op           45 B/op          0 allocs/op
```

