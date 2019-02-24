## Week 3 lecture notes

We will be primarily looking at composite datatypes in this week.

### Arrays

- Fixed length of series of elements of same type.
- Elements indexed using subscript notation `[]`. Index starts at `0`
- Elements are initialized to zero value

```
var x [5]int
x[2] =2 
fmt.Printf(x[1]) // prints 0
```

- Array Literal: Predifined set of values that goes into an array.

```
var x [5]int = [5]{1,2,3,4,5}
```

can use `...` as a shorthand for array literals.

```
x := [...]int{1,2,3,4}
```

- Iterating through Array using `for`. `range` returns two values index and element at index.

```
x:= [3]int {1,2,3}

for i,v range x {
    fmt.Printf("ind => %d value => %d",i,v)
}
```

### Slices

- A "window" on an underlying array and can have variable size upto the size of the underlying array.

A slice usually have the below attributes

- Pointer which indicates the start of the slice.
- Length which indicates the number of elements in the slice.
- Capacity is the maximum number of elements in the slice. It would essentially be the number of elements possible from the start of the slice to the end of the underlying array

```
arr := [...]string {"a","b","c","d","e"}

s1 := arr[1:3] //slice containing elements at index 1 and 2
s2 := arr[2:4] // slices can have overlaps.
```
- `len()` returns length of slice.
- `cap()` returns the capacity of slice.
- Writing to slices changes the underlying array as well.
- Slice literals can be used to initialize a slice. It can be defined as below
```
sli := []int{1,2,3} // We didn't put anything under bracked.
```
What happens here is , go creates the underlying array of the same size and then create a slice with pointer starting at `0` and capacity set as the size of the underlying array.
The compiler identifies this as a slice literal from the empty square bracket.

### Variable Slices

- Slice can be also created using `make()`. 
- Initializes elements to zero by default.

2 argument version, type of the array and length/capacity.

```
sli = make([]int , 10)
```
3 argument version : type, lenth, capacity(underlying array)

```
sli = make([]int, 10, 15)
```

- `append()` can be used to append to the end of slice. It increases the size of underlying array if necessary, but we don't want to do that because of a time penality.
```
sli = make([]int, 0, 3)
sli = append(sli, 100)
```

### Hashtable

- Key, value pair
- hash fn is used to calculate the key and provide the slot for the key.
- Map is golang's implementation of hashtable.
- we use `make()` to create maps.

```
var idMap map[string]int //Declaration , here string is the key and int is the value
idMap = make(map[string]int)
```
- Map literal
```
idMap:= map[string]int {
    "joe" : 123
}
fmt.Printlln(idMap["joe"])

idMap["jane"] = 456
delete(idMap, "joe")
```

- Two value assignment tests for existence of a key.

```
id, p := idMap["joe"]
```

id is value, p is a boolean for the presence of key.

- `len()` returns number of values.

- Iterating through the map.

```
for key,value := range idMap{
    fmt.Println(key, value)
}
```

### Structs

- Aggregate data types.
- Groups together other objects of other data types

```
type struct Person{
    name string
    addr string
    phone string
}

var p1 Person
```
- Accessing using dot notation.

``` 
p1.addr
```
- Initialization is done using `new()` and initializes the fields to 0 as below.

```
p1 := new(Person)
```

Initialization with literals can be as below

```
p1 :=  Person(name:"joe",
              addr:"address",
              phone:"1234")
```
