package main

// use parentheses to import multiple
import (
	"fmt"
	"math"

	"github.com/bushki/go_course/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.71))
	fmt.Println(math.Ceil(2.71))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("hello"))
}
