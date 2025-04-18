##
### Performance in the cloud:
- we've made a deliberate choice to accept some overhead
Trade off performance against other things like:
    - choice of architecture
    - quality, reliability, scalability
    - cost of development & ownership

-> but I still want simplicity, readability & maintainability of my code.

- when thinking of optimization we should think about it from top top button: 
Architecture: latency, cost of communication
Design: algo, concurrency, layers
Implementation : prog languages, memory use
=> Mechanical sympathy play a role in my implementation.
Fucking interpreted languages my cost 10x more to operate due to their inefficiency. Because it abstract the machine for the developer 

memory performance: [cpu]---d---[dram]: even the cpu is becoming more faster but the distance between the cpu and dram it didn't really changed. => access latency (remember clock cycle that a cpu have to make in order to access a memory)

cpu<-- l1-cache<-- l2-cache<--l3-cache(shared cache)<-- memory 64byte
* even thought when have multiple core (4 core) in each core we have three cache level, there is a need to synchronize cache between all cores.

Cache: locality in space/time
* Things that make the cache less efficient:
    * synchronization between Cpus
    * copying blocks of data around in memory
    * non-sequential access patterns(calling func, chasing pointers)

    a little copying is better than a lot of pointer chasing!

* Things that make the cache more efficient:
    * keeping code or data in cache longer
    * keeping data together (so all of a cache line is used)
    * processing memory in sequential order (code or data)


----------------------------------------------------------------

Mechanical sympathy: 
* make software simpler.
* make software that works with the machine, not against it.
make software suck less

* 
* Access patterns:
    *A slice of objects beats a list with pointers
    * A struct with contiguous fields beats a class with pointers
    * Calling lots short methods via dynamic dispatch is very expensive


* Synchronization costs
synchronization has two cost:
    * the actual cost to synchronize(lock&unlock)
    * the impact of contention if we create a hot spot (lots of goroutine try to access the same mutex )

* Other costs
    * disk access
    * garbage collection
    * virtual memory & its cache
    * context switching between processes


-------------------------------------------------------
*Optimization
• to allocate contiguously
• to copy or not copy
• to allocate on the stack or heap (sometimes)
• to be synchronous or asynchronous
• to avoid unnecessary abstraction layers
• to avoid short / fowarding method


There are only three optimizations:
1. Do less
2. Do it less often
3. Do it faster
The largest gains come from 1, but we spend all our time on 3.”
— Michael Fromberger
