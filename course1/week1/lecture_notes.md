## Week 1 Lecture notes

### Advantages of Go

- Concurency inbuilt.
- Faster run time.
- Simple OOP constructs when compared to likes of C++
- Garbage collection

### Software Translation

- ideas on machine code, assembly code , high level languages (HLL's) etc.
- Go has garbage collection inbuilt, it is an efficient one and hence doesn't slow
down much. 


### Object Oriented Programming

- Go doesn't use the keyword `class` , but instead use `struct`
- Structs don't have inheritance, constructors, generics.


### Concurrency 

- `Goroutines` represent concurrent task.
- `channels` are used to communicate between tasks.
- `select` enables task synchronization.


### Go Tool

- `go build` compiles the program.
- `go doc` prints documentation for the project.
- `go fmt`  formats the code.
- `go get` installs the packages.
- `go test` run tests.


###  Variables
- usual rules
- variable should have a name and type.
  eg: var x int , var x,y int
- can create aliases for types.
- uninitialized variables have value set as zero.
- Short hand for variable declaration and initialization is to use `:=` .
  It can be done inside function, not outside.

