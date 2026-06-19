package main

import (
	"fmt"
)
func main(){
	fmt.Println(add(43,13))
	fmt.Println(subtract(4,2))
     a,b:=swap("hello","world")
     fmt.Println(a,b)
	 fmt.Println(split(17))
}

func add(x int,y int)int{
	return x+y
}
func subtract(x,y int)int{
	return x-y
}
var c,python,java bool
//var statement declares al list of variables as in function 
//argument list the type is last.

//multiple results:
//A function can return any number of results
//Illustration
func swap(x,y string)(string,string){
	return y,x

}

func split(sum int)(x,y int){
	x=sum*4/9
	y=sum-x
	return
}

