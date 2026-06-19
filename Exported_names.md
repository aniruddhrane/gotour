In go an entity(name) is exported if it begins with camel case or upper case.
For example Pizza is an exported name as is Pi which is exported from the math package.
package main

import(
    "fmt"
    "math"
)
func main()
{
    fmt.Println(math.pi)  //error as no pi was exported because it was never allowed
}