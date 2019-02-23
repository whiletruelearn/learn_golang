## Week 2 lecture notes

### Pointers

- A pointer variable holds the address of data in memory
- `&` operator returns the address of a variable/function.
- `*` operator returns data at an address (dereferencing)

Example:

```
var x int = 1
var y int
var ip *int // ip pointer to int
ip = &x  // ip now points to x variable
y = *ip // y is now 1
```

- `New` is an alternate way to create a variable.
- `new()` creates a variable and returns a pointer to the variable and 
variable is initialized to zero by default.

```
ptr := new(int) // ptr holds the address of new variable.
*ptr = 3  // assign 3 directly to the address location
```

### Variable Scope

- Global and local scopes.
- Scope is defined using blocks . Blocks are defined by `{}`
- Hierarchy of blocks, File block, Package block . 
- Implicit blocks viz `for, if, switch` etc.


### Deallocation memory

- Deallocate objects when they are no longer needed.
- `Stack` is an area of memory dedicated to function calls.
- Local variables for a function are traditionally allocated in stack and they are deallocated when the function completes.
- `Heap` on the other hand is a persistent memory , it needs to be explicitly deallocated.
   eg: global variables.

```
#C : Explicit deallocation 
x= malloc(32);
free(x);
```

- This could be error prone, but unline in an intepreted language which does deallocation
by itself, this could be fast.

### Deallocation is tricky 

- It is hard to determine when a variable is no longer use. Especially
when we use pointers

```
func foo() * int {
    x := 1
    return &x
}

func main() {
    var y *int
    y = foo()
    fmt.Printf("%d", *y)
}
```

- Calling `foo` returns pointer to `x`. Ideally we would want to deallocate a variable when 
function call is complete. But here we need to hold on to `x` since the scope is alive in `main`.
- This is why some of the modern HLL's have completely removed pointers. it is a design choice
made such that their automatic garbage collection systems doesn't get screwed up.
- Garbage collectors in python for instance use something called reference count to see a variable should be deallocated. In Cpython , every variable would be having 
- Go is a compiled language with garbage collector built in.


### Comments , Printing, Integers

- `//` => comment
- `/* */` => Block comment
- `fmt.Printf()` does printing
- `int8, int16, int32, int64, uint8, uint16, unint32, uint64`
-  Arithmetic operators : `+ - * / % << >>`
-  Boolean operators : `|| &&`


### Ints, Floats, Strings

- Type conversion is done by specifying the `T()` operation

```
var x int int8 = 1
var y = int16(x)
```

- `ASCII` is 8 bit code while `unicode` is 32 bit code.
- `utf-8` is a variable length coding system, it can be 8 bit or it can go upto 32 bit.
Go uses utf-8 code points.
- Strings are read only and can't be modified. 


### String packages

- `Atoi`, `ItoA` etc

### constants

- `const` is used to define constants whose values doesn't change.
- `iota` is a set of related but distinct constants. eg: days of week. Similar to enum in python.

```
type Grades int
const (
    A Grades = iota
    B
    C
    D
    F
)
```
Here each `Grades` constant would be assigned a unique integer value.

### Control Structures

- `if` 
- `for`
```
//format 1

for i:=0; i <10; i++ {
    fmt.Printf("hi)
} 

//format 2

i = 0
for i <10 {
    fmt.Printf("hi")
    i++
}

//format 3

for {
    fmt.Printf("hi")
}
```


- `switch` is a multi-way if statement.
```
switch x {
    case 1:
        fmt.Printf("case 1")
    case 2:
        fmt.Printf("case 2")
    default:
        fmt.Printf("no case")
}
```

- Tagless `switch` can be used instead of multiple if else statements.
```
switch 
{
    case x > 1:
        fmt.Printf("case 1")
    case x < -1:
        fmt.Printf("case 2")
    default:
        fmt.Printf("no case")

}
```
The first expression that is evaluated to be true is executed.

- `break` exists the containing loop when the boolean condition is met.
- `continue` skips the iteration
- `scan` reads user input. Takes a pointer as an argument.
   Returns number of scanned items and error codes

```
var appleNum int
fmt.Printf("Number of apples?")
num, err := fmt.Scan(&appleNum)
fmt.Printf(appleNum)
```