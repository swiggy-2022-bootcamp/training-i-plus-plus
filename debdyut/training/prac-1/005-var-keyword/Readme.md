

### short declaration cannot be used outside function

```go
package main

import "fmt"

x := 42

func main () {	
	fmt.Println(x)
}
```

<strong>Output</strong>

```bash
# command-line-arguments
.\main.go:5:1: syntax error: non-declaration statement outside function body
```