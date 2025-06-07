# DSA
## Basic data structures
A named location that can be used to store and organize data.
Example Array (a collection of elements stored of contiguous memory location), Tree (A hierarchy of family relationships)

Algorithm: a collection of steps to solve a problem.

Why ? To write code that is both time and memory efficient. && for job interviews.

#### Stack
LIFO Last-In First-out.
Store objects into a sort of "vertical tower".
push(): to add to the top
pop(): to remove from the top it return the poped object ///excepiton EmptyStackException if the stack already empty 
peek(): get the top of the stack
empty() : return a boolean 
search(): return 1 or -1

* Usage of stack in real word:
    Undo/redo features in text editors
    Moving back/forword throught browser history
    Backtracking algorithms (maze, file directories)
    Calling functions (call stack)

[see DSA/Stack]

#### Queue
FIFO: First-In First-Out example (line of people)
A collection designed for holding elements prior to processing
Linear data structure 

add() =  offer()
remove() =  poll()
element() = peek()
Note: based on docs use the second column(poll doesn't cause an exception)

Queue interface extend Collection so it has the same method of collection: isEmpty, size(), contains(value)

* Usage of queue in real word:
    keyboard Buffer (letters should appear on the screen in the order they're pressed)
    Printer Queue (Print jobs should be completed in order)
    Used in linkedList, PriorityList, Breadth-first search

#### Priority Queue
FIFO data structure that serves elements with the highest priorities first before elements with lower priority.
It's like Queue but it sort it's elements before dequeuing them.

#### LinkedLists (singly)
Array and ArrayList store elements in contiguous memory location.
  0    1   2                       8
["M" |"A"|"H"|"M"|"O"|"U"|"D"|" "|"!"||]
@13x @15x                           @23x
are good for randomly accessing elements because they have index, but are not good for inserting or deleting elements.
for example if I need to insert the letter "M" at the index 3 , I need  first to shift the rest of the array and then put the letter in the right position.(think of inserting many letters) 
and the same for deletion.
===> LinkedList have the advantage(better a insertion and deletion)

LinkList made of long list of nodes, each node contain two parts the data to store and the address *Pointer* of the next node in line. non contiguous memory locations.

["A"|@x1259F]---> ["B"|@x1h98]---> ["C"|@x1259]---> ["D"|null]
* insertion : O(1)!:!
["A"|@x1259F]                      ["B"|@x1h98]---> ["C"|@x1259]---> ["D"|null]
             ---> ["$"|@x1259F]--->

["A"|@x1259F]--->["$"|@x1259F]--->["B"|@x1h98]---> ["C"|@x1259]---> ["D"|null]
* Deletion: O(1)
["A"|@x1259F]--->--->["B"|@x1h98]---> ["C"|@x1259]---> ["D"|null]

! Bad at searching (we don't have indexes like arrays) To find an element we need to begin from the head to the tail until we found the
element.(Linear time O(n))


#### LinkedLists (doubly)

[null|"A"|@x1259F]---> [@x1259F|"B"|@x1h98]---> [@x1h98|"C"|@x1259]---> [@x1259|"D"|null]

Benefits of doubly LinkedList:  
    is that we can traverse the from head to tail or from tail to head.
    insertion and deletion of nodes is O(1).
    No/LOW memory waste
Drawbacks :
    uses more memory than singly LinkedList
    No random access of elements(no indexing)
    Accessing and searching element is O(n)


LinkedList class implements List and Deque Interfaces many methods like[addFirst,getFirst,removeFirst, pollFirst .... see dcocs,] and also we can treat it as stack or a queue.

Usage: 
    implement Stack/Queues
    GPS navigation
    music playlist

#### Array
Array: fixed capacity
Dynamic array : java ArrayList, Go Slice, C++ Vector, JavaScript Array, Python List.
when the inner static array reach the limited capacity , Dynamic array will intiate and create new array (java->1.5, go->2,rust->2, pyton->1.125) and then  copy old value to the new array.

Accessing an elements: O(1)
Searching: O(n)
Insertion or deletion: O(n) // but insert/delete at the end is easy

Disadvantages:
    waste more memory (more than linkedList)
    shifting elements is time consuming
    expanding/shrinking the array is time consuming

--------------------------------------------------

--------------------------------------------------
## Big O notation
How code slows as data grows.
Describes the performance of an algorithm as the amount of data increases.
Machine independent (nbre of steps to complete an algorithm)
Ignore smaller operation o(n+1) -> o(n)


O(1): constant time
    random access of an element in an array
    inserting at the beginning of linkedList

O(n): linear time
    looping through elements in an array
    searching through a linkedList

O(log n): logarithmic time (the more the data increases the more the algorithm become more efficient.
    binary search 

O(n log n): quasilinear time 
    quicksort
    mergesort
    heapsort

O(n^2): quadratic time
    insertion sort 
    selection sort
    bubblesort

O(n!): factorial time
    Traveling Salesman Problem

n: amount of data









--------------------------------------------------
## Searching algorithms
* Linear search: one by one, examine the elements of an array to find a value.
* Binary search: finds the position of a target value within a sorted array. Half of the array is eliminated during each step.
O(log n)























--------------------------------------------------
## Sorting algorithms
--------------------------------------------------
## Graphs
--------------------------------------------------
## Trees
--------------------------------------------------
