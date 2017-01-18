package samples

import (
	"fmt"
	"github.com/kmkatsma/hello/models"
	"math"
)

func createVertex(n int) {
	v := models.Vertex{X: 1, Y: 2}
	v.X = n
}

func counting() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

func arrays() {
	var a [2]string
	a[0] = "Hello"
}

func pointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
