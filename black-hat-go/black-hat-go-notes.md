# Black Hat Go

### Go Tool Commands

**Run a go main package:**

`go run main.go`

**Build a go binary:**

`go build main.go`

**Build a stripped go binary:**

`go build -ldflags "-w -s" main.go`

**Cross-compiling (supported arch https://golang.org/doc/install/source#environment/):**

`GOOS="linux" GOARCH="amd64" go build main.go`

**The go doc command:**

`go doc library.function`

Displays documentation for a given function

**The go get command:**

Requires the execution of `go mod init` in the project dir as of version 1.16

`go get github.com/user/package-repo-name`

**The go fmt command:**

`go fmt main.go`

Formats your source code- most IDEs do this on save

**The golint and go vet commands:**

golint is a standalone tool, you’ll need to install it separately by using 
`go get -u golang.org/x/lint/golint.`

`go vet` inspects your code and uses heuristics to identify issues

**Go Playground:**

Execution environment - https://play.golang.org/

**Other commands and tools:**

- `go test` to test projects
- `go tool cover` to check test coverage
- `goimports` to fix imports

### Understanding Go Syntax

**Data Types:**

Primitive:

- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- uintptr
- byte
- rune
- float32
- float64
- complex64
- complex128

**Declare variables:**

```
var x = "Hello, World!"
z := int(42)
```

**Slices and Maps:**

Slices are like arrays that you can dynamically resize and pass to functions more efficiently.

Maps are associative arrays, unordered lists of key/value pairs that allow you to efficiently 
and quickly look up values for a unique key.

slice and map declaration:

```
var slice = make([], string, 0)
var map = make(map[string]string)

slice = append(slice, "a string")
map["a key"] = "a value"
```

**Pointers, Structs, and Interfaces:**

Go uses the same syntax as C for pointers, &varname retrieves the address in memory of the 
variable, and the * operator to dereference the address

```
var count = int(42)
var ptr int* = &count
fmt.Println(*ptr)
*ptr = 100
fmt.Println(count)
```

New data types are declared using struct

```
type Person struct{
    Name string
}
```

Interfaces define a set of actions that must be implemented for any concrete implementation to 
be considered a type of that interface

type Friend interface {
    SayHello()
}

**Control Structures**

if/else statements

```
if x == 1 {
    fmt.Println("X is equal to 1")
} else {
    fmt.Println("X is not equal to 1")
}
```

switch statements

switch x {
    case "foo":
        fmt.Println("Found foo")
    case "baz":
        fmt.Println("Found baz")
    default:
        fmt.Println("Default case")
}

Cases don't require the use of break statements

type switch

```
func foo(i interface{}) {
    switch v := i.(type)v {
    case int:
        fmt.Println("I'm an integer!", v)
    case string:
        fmt.Println("I'm a string!", v)
    default:
        fmt.Println("Unknown type!", v)
    }
}
```

Uses special syntax to compare the type of the interface passed in

for loops

```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

using range with collections (slices, maps):

```
nums = []int{1,2,3,4,5}
for idx, val := range nums {
    fmt.Println(idx, val)
}
```

**Concurrency**


