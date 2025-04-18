
*  Concurrency: Parts of the program may execute independently in some non-deterministic order.(can have it wit a single-core processor)
is about dealing with things happening out-of-order.
* Parallelism : Parts of a program execute independently at the same time.(happen only on multi-core processor)
is about dealing with things happening at the same time.
A single program won't have parallelism without concurrency
=> Concurrency doesn't make the program faster, parallelism does.

*Race condition:
- non-deterministic program may produce invalid results
(2 can be first or 3 can be first)
 1. {1, 2a, 2b, 3a, 3b, 4} correct
 2. {1, 2a, 3a, 2b, 3b, 4} incorrect
 3. {1, 2a, 3a, 3b, 2b, 4} incorrect
 4. {1, 3a, 3b, 2a, 2b, 4} correct
 5. {1, 3a, 2a, 3b, 2b, 4} incorrect

- changing things that are shared
remember Two deposits to a bank account example.
1. Solution [(read, modify, write)] lock them together => make them atomic operation (non divisible).
2. don't share anything
3. make the shared things read-only
3. allow only one writer to the shared things


* Channel: is one-way communications pipe (unix pipe)
things go in one end, come out the other
in the same order they went in
until the channel is closed
multiple readers & writers can share it safely

		°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°
--->																				--->
		°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°
- a method of synchronization as well as communication
we know that a send(write) always happens before a receive (read)


* CSP: Communicating sequential processes

* Goroutine: is a unit of independent execution (coroutine) go in front of a func call. => start an independent thread*
-> need to be stopped

