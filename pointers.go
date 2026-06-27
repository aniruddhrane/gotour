package main
import "fmt"
//variable is stored somewhere in memory it has name 
//it is a box with value
//we give this box an address to access or locate that 
//somewhere
func main(){
//var foo int =23
i,j:=42,2701
fmt.Println(&i,&j)
p:=&i //here see this is the declaration of the pointer
fmt.Println(*p)
p=&j
*p=*p/37
fmt.Println(j)
/*
 the * is used in two types of scenarios
 *int this whole thing is a type:pointer type with int as a base
 *p:* acts as a operator returns what p is pointing to 

*/
/* 
   memory allocations
   goroutines
*/
}