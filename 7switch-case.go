package main

import "fmt"

func main() {
	i := 3
	switch i {
	case 1:
		fmt.Print("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3") // 3 4
		fallthrough
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println(" a number")
	}


}
