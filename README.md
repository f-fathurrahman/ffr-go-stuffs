Learning Go

Compile and link the program:

```
go build program_name.go
```

This will produce executable `program_name`

The compiler will complain and will not produce executable if there are any
unused variables.

Using `gofmt` to reformat the file:
```
gofmt -w filename.go
```