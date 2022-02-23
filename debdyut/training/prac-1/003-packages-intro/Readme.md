
### Hello world program

Uses fmt (format, pronounced as 'fumpt') package.

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

<strong>Output<strong/>

```bash
Hello, World!
```

### fmt.Println return types

`func Println(a ...interface{}) (n int, err error)` <br/>
Println formats using the default formats for its operands and writes to standard output. Spaces are always added between operands and a newline is appended. It returns the number of bytes written and any write error encountered.

```go
func main() {
	n, err := fmt.Println("Hello, World!")
	fmt.Println("Bytes written:", n)
	fmt.Println("Error encountered:", err)
}
```

<strong>Output<strong/>

```bash
Hello, World!
Bytes written: 14
Error encountered: <nil>
```

### Must use all varaibles (avoid code pollution)

```go
func main() {
	n, err := fmt.Println("Hello, World!")
	fmt.Println("Bytes written:", n)
	// fmt.Println("Error encountered:", err)
}
```

```bash
# command-line-arguments
.\main.go:16:5: err declared but not used
```

### Skip variables from return

The above issue can be resolved by using `_` to avoid declaring `err`.

```go
func main() {
	n, _ := fmt.Println("Hello, World!")
	fmt.Println("Bytes written:", n)
}
```

```bash
Hello, World!
Bytes written: 14
```