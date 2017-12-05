package main

import "fmt"

func main() {
	n := 5
		//init, condition, post; exactly like for loop
	for i := 0; i < n; i++ {
		j := 10
			//condition; similar to while-do.
		for j > 0 {
			fmt.Println("j :", j)
			j--
			k := 0
				//infinite loop
			for {
				if k == 3 {
					break
				} else {
					fmt.Println("\t\t k ",k)
					k++
				}
			}
		}
		fmt.Println("\t i ", i)

	}
}
