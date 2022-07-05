/*
Package deque provides a fast double ended queue implementation.

Deque allows the addition & removal of items at either of the queue
in O(1) perfomance. FIFO and LIFO operations are both supported.
The internal nodes are abstracted from the user and the stored value is only exposed.

Benchmarking

Benchmark Tests that do not contain a suffix use the Deque.
Tests ending in `_List` use the standard slice.
```
BenchmarkAppendFirst-16                 20438316                65.57 ns/op           24 B/op          1 allocs/op
BenchmarkAppendFirst_List-16            100000000               30.27 ns/op           45 B/op          0 allocs/op

BenchmarkAppendPopFirst-16              17881148                78.14 ns/op           24 B/op          1 allocs/op
BenchmarkAppendPopFirst_List-16         100000000               32.96 ns/op           45 B/op          0 allocs/op

BenchmarkInsertAt-16                    20175003                66.46 ns/op           24 B/op          1 allocs/op
BenchmarkInsertAt_List-16               100000000               34.19 ns/op           45 B/op          0 allocs/op
```

Generics

This deque is allowed to be made up of any value the user specifies.

Empty Deque

This queue implementation returns a boolean as a second value in certain operations.
True means the first value returned is part of the queue.
False means the first value returned did not exist. The value returned is the zeroth value of the type.
*/
package deque
