/*
Package deque provides a fast double ended queue implementation.

Deque allows the addition & removal of items at either of the queue
in O(1) perfomance. FIFO and LIFO operations are both supported.
The internal nodes are abstracted from the user and the stored value is only exposed.

Benchmarking

To be added.

Generics

This deque is allowed to be made up of any value the user specifies.

Empty Deque

This queue implementation returns a boolean as a second value in certain operations.
True means the first value returned is part of the queue.
False means the first value returned did not exist. The value returned is the zeroth value of the type.
*/
package deque
