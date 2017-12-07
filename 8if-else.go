package main

import (
	"os"
	"fmt"
)

func main(){
	if err:= FileOpen...; err!=nil{	//initialization in if statement, scope is limited
		fmt.Println("error")
	} else {
		fmt.Println("do something")
	}
}
