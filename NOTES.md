## Hello World

```go
// a comment
/*
Multiline
comment
*/
import "fmt"
func main() {
  fmt.Printf("Hello Go from ffr\n")
}
```

Explore package `fmt`.

- `fmt.Println`
- `fmt.Print`
- `fmt.Fprintf`: write to file


## Declaring variable

```go
var i int = 1
var a float64 = 2.3
```

Using type inference:
```go
i := 1
a := 2.2
```

Variable declared using `:=` cannot be redefined using `:=`, we need to use `=`
instead.
```go
i := 0
i := 1 // error
i = i + 1 // OK
```


## Functions

```go
func sayHello() {
  fmt.Printf("Hello from ffr\n")
}

func getName() string {
  return "efefer"
}
```
