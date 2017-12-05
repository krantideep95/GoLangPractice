package main

import "fmt"

func main() {
	a := 5
	b:= &a
	fmt.Println("a ", a, " address ", &a)
	initi(b)
	fmt.Println(a," =a, *b= ",*b)
}

func initi(x *int) {
	*x = 0
}
