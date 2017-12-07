package main

import (

	"fmt"
)

func main() {

	a := 4
	b := 5
	AddnPrint(a, b)
	fmt.Println(avg(a, b))
}

func AddnPrint(x, y int) {
	fmt.Println(x + y)
}

func avg(arg ...int) (s float32) {

	for _, ar := range arg {
		s += float32(ar)

	}
	s /= float32(len(arg))
	return

}
