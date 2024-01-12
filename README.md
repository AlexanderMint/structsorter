Fork: [fieldalignment](https://cs.opensource.google/go/x/tools/+/master:go/analysis/passes/fieldalignment/fieldalignment.go)

Added sorting of fields in alphabetical order when the size of the fields is the same

### Install

```bash
go install github.com/AlexanderMint/structsorter/cmd/structsorter
```

### Example

```go
// test/in_file.go
type A struct {
	cc *int
	b  string
	i  int
	a  string
	z  int64
	aa bool
	c  string
	bb *int
}
```

```bash
structsorter --fix ./test/in_file.go
```

result:

```go
type A struct {
    bb *int
    cc *int
    a  string
    b  string
    c  string
    i  int
    z  int64
    aa bool
}
```