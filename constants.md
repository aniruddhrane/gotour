Constants are declared like variables but with const keyword.

They can be character,string,boolean or numeric values

They cannot be declared using :=

ex.
package main
import(
  "fmt"
)

const Pi=3.14

func main(){
     const World="earth"
     fmt.Println("Hello",World)
     fmt.Println("Happy",Pi,"Day")
     
     const Truth=true
     fmt.Println("Go rules?",Truth)
}
