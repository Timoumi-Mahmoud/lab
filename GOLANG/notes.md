## 

----------------------------
### Strings
- The internal string representation is a pointer and a length:
- are immutable, and can share the underlying storage.
- Are all unicode (technique to represent special character from different languages) replace ASCI (uint8)
- byte : synonym for uint8
- rune : a synonym for int32 for characters --> which means 4 bytes (but it's a waste of resources ) -> UTF-8 encoding is an efficient way to encode unicode.


- string to rune
```
s := "èlite" ==> length is 6 not 5
[]rune(s)  // []int32 [232 108 105 116 101]
[]byte(s)  // []int32 [192 159 232 108 105 116 101]  ==> unicode in action here it represent è by two bytes

```
Length of a string is the numbers of bytes string required to represent the unicode character  ==> physically 

- convert string to []string : 
```
sliceOfString = []string{"Mahmoud"}
```
-  conert slice of strings to strings 
```
names =[]string{"bara", "khadija", "salim"}
strings.Join(names, ", ")
```
-------------------------
### Arrays: 
- Fixed size which is fixed at compile time: var arr [4]int
- Mutable
- are passed by value, thus elements are copied 

-------------------------
### Slices: 
- have variable length, backed by some array
- they are passed by referencce; no copying , updating OK.
- Pointer | cap | length ==> slice discripter


slice				vs					Array
- Variable length			- Length fixed at complite time
- Passed by reference		- Passed by value(copied)
- Not comparable			- Comparable (==)
- Cannot be used as map key - can be used as map key
- Has Copy & append methods

-----------------------------
### Maps:
- Maps are dictionaries: indexed by key, returning a value.
- can read from a nil map, but inserting will panic.
- Are passed by reference; no copying, updating OK.
- The type for the key must have == and != defined (not slices, maps , or funcs)
- not ordered : => to sort a map: extract the keys -> sort them --> loop over the map using the sorted keys

------------------------------
# File I/O
- Package os has functions to open or create files , list diretories ,etc. and hosts the file type.
- Package io has utilities to read and write.
- package bufio provides the buffered I/O scanners etc.
- package io/ioutil has extra utilities such as reading an entire file to memory , or writing it out all at once.
- Package strconv has utilities to convert to /from string representations. 



-------------------------------
## Func
- formal parameters and actual parameters.
- A parameter is passed by value if the function gets a copy; the caller can't see changes to the copy.
	. numbers
	. bool
	. arrays
	. structs !! uness a pointer is used (example a.copies++ see album example)
- A parameter is passed by reference if the function can modify the actual parameter such that the caller sees the changes.
	. things passeed by pointer(&x)
	. strings (but they're immutable)
	. slices
	. maps
	. channels

-------------------------------------
## Defer:
- close a file we opened
- close a socket /http request we made
- unlock a mutex we locked
- makesure something gets saved before we're done
-------------------------------------
# Closure
- function uses variables from another function scope( outer function)

