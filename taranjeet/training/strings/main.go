package strings

import (
	"fmt"
	"strings"
)

func StringOperation(a, b string) {
	fmt.Println(strings.Compare(a, b))
	fmt.Println(strings.EqualFold(a, b))
	fmt.Println(strings.Contains(a, b))
	fmt.Println(strings.ContainsAny(a, b))
	fmt.Println(strings.Index(a, b))
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.Trim(a, "/"))
	fmt.Println(strings.TrimSpace(a))
}
