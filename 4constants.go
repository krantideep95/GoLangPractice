package main

import (
	"fmt"
)

const cb = 23

func main() {
	const a = 5
	fmt.Println("a", a, "b", cb)
	const q = iota
	const t float32 = iota
	const r


	fmt.Println(q, " t ", t)
}
