//Go has only one looping construct, the for loop
//The for loot have tree components
//init statement:execute before the iteration
//the condition expression
//the post statement exectued after every iteration

package main
import "fmt"

func main(){
	sum:=0
	for i:=0;i<10;i++{
		sum+=i
	}
	fmt.Println(sum)

	//the init and post statements are optional
    sum=1
	for ;sum<1000;{
		sum+=sum
	}
	fmt.Println(sum)
	//while mimicry
	sum=1
	for sum<1000{
		sum+=sum
	}
	fmt.Println(sum)
	//for{
      //this is infinite loop with no condition
	//}
}
