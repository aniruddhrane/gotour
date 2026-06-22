package main

import(
	"fmt"
	"math"
)

func sqrt(x float64) string{
	if x<0{
		return sqrt(-x)+"i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func main(){
	fmt.Println(sqrt(2),sqrt(-4))
}
//go if statement do not require parenthesis but only braces
//{}
//like a for loop we can intitialize a variable in 
//if statement itself
//if v:=math.Pow(x,n);v<lim{
// return v
//}
