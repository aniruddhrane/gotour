package main
//A struct is a collection of fields
import "fmt"

type Vertex struct{
	X int
	Y int
}

func main(){
	fmt.Println(Vertex{1,2})
	v:=Vertex{1,2}
	v.X=4
	fmt.Println(v.X)
	ver:=Vertex(1,2)
	p:=&ver
	p.X=1e9
	fmt.Println(v)
}