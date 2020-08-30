# GO NOTES
Quick Notes for the Go Lang

#### Shortening 
`x int, y int => x, y int`

#### Multiple Results
`func swap(x, y string) (string, string)`

##### Data Type
`bool`
`string`
`int  int8  int16  int32  int64`
`uint uint8 uint16 uint32 uint64 uintptr`
`byte // alias for uint8`
`float32 float64`
`complex64 complex128`

#### Type Conversion
The expression T(v) converts the value v to the type T
`i := 42`
`f := float64(i)`
`u := uint(f)`

#### Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
`defer fmt.Println("world")`
`fmt.Println("hello")`
would print `hello world`

#### Pointers
Go has pointers. A pointer holds the memory address of a value.
The type  `*T`  is a pointer to a  `T`  value. Its zero value is  `nil`.
`var p *int`
The  `&`  operator generates a pointer to its operand.
`i := 42`
`p = &i`
The  `*`  operator denotes the pointer's underlying value.
`fmt.Println(*p) // read i through the pointer p`
`*p = 21         // set i through the pointer p`
This is known as "dereferencing" or "indirecting".
Unlike C, Go has no pointer arithmetic.
Example - 
	
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

#### Arrays

The type  `[n]T`  is an array of  `n`  values of type  `T`.
The expression
`var a [10]int`
declares a variable  `a`  as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.

## Mutating Maps

Insert or update an element in map  `m`:
`m[key] = elem`
Retrieve an element:
`elem = m[key]`
Delete an element:
`delete(m, key)`
Test that a key is present with a two-value assignment:
`elem, ok = m[key]`
If  `key`  is in  `m`,  `ok`  is  `true`. If not,  `ok`  is  `false`.
If  `key`  is not in the map, then  `elem`  is the zero value for the map's element type.
**Note:**  If  `elem`  or  `ok`  have not yet been declared you could use a short declaration form:
elem, ok := m[key]

#### Notes 
- Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
- The zero value is:
	-   `0`  for numeric types,
	-   `false`  for the boolean type, and
	-   `""`  (the empty string) for strings.
- `%T`for printing type -> `fmt.Printf("v is of type %T\n", v)`
- `if v := math.Pow(x, n); v < lim` If statements can have init statements just like `for loop`. Scope is limited to `if-else`block
- The type `[]T` is a slice with elements of type `T`.
- A slice does not store any data, it just describes a section of an underlying array. Changing the elements of a slice modifies the corresponding elements of its underlying array. Other slices that share the same underlying array will see those changes.
