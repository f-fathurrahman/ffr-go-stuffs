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

## Accessing documentation offline

```
godoc -htpp=:6060
```

Use browser to access `http://localhost:6060`.


## Declaring variable

```go
var i int = 1
var a float64 = 2.3
var is_working bool = true
```

Using type inference:
```go
i := 1
a := 2.2
```

Declare multiple variables:
```go
name, age := "Fadjar F", 29
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
