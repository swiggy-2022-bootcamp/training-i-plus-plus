
import (  
	"fmt"
)

func hello(i int) {  
	fmt.Println(i)
}
func main() {  
	i := 5
	defer hello(i)
	i = i + 10
}