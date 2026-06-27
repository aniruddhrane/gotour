package main

import(
	"fmt"
)
func main(){
	fmt.Println("counting")
	for i:=0;i<10;i++{
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
//deferring function calls are pushed onto a stack
//when a function reuturns its deffered calls are 
//executed in last in first out order