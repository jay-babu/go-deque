/*
Package deque provides a fast double ended queue implementation.

Deque allows the addition & removal of items at either of the queue. FIFO and LIFO operations are both supported.
The internal nodes are abstracted from the user and the stored value is only exposed.

This library is mostly a wrapper around a list with some logic abstracted away.
If an operation would throw an error on a list, it will be thrown here.

> Implementation using a traditional deque was done with O(1) operations
at the beginning and end. Using a list turned out to be more performant.

Benchmark

To be added.

Generics

This deque is allowed to be made up of any value the user specifies.

*/
package deque
