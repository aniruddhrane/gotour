package main
import(
	"fmt"
	"os"
)

func main(){
  data:=[]byte("Hello Go!")
  
  err:=os.WriteFile("hello.txt",data,0644)
  //files dont understand strings 
  //files understand bytes
  if err!=nil{
	fmt.Println(err)
	return
  }
  fmt.Println("File written successfullly")
}
// Its signature is approximately

// func WriteFile(
//     filename string,
//     data []byte,
//     perm FileMode,
// ) error

// Notice
// It returns only
// error
// because there is nothing meaningful to return on success.
// First parameter
// "hello.txt"
// File name.
// Second parameter
// data
// Bytes to write.
// Third parameter
// 0644
// This surprises everyone.
// It is the file permission.
// For now, simply remember:
// 0644
// is the standard value for normal text files.
