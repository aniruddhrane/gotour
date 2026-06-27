package main

//unlike switch in other language go does not require
//break statements  the first case gets executed with 
//the condition matching block
//go switch is shorter way of writing if-else statements

import(
	"fmt"
	"runtime"
)
func main(){
switch os:=runtime.GOOS; 
//runtime.GOOS is the constant that defines the target 
//operating system(e.g darwin(macos),linux,windows)
//the switch statement intitializes the local variable os
//before evaluating it
//implicit break:Go automatically break after matching 
//a case so no break statements are required
os{
case "darwin":
	fmt.Println("macOS")
case "linux":
	fmt.Println("Linux")
default:
	fmt.Println("%s.\n",os)
}
}


//runtime library is a collection of built in software routines
//that an operating system uses to execute a program while its running
// It acts as a bridge between your written code, 
// the programming language's internal features, and the underlying computer hardware.
//Key FunctionsMemory Management: Allocates and frees RAM (like garbage collection in Go or Java).
// Thread Management: Handles concurrency, managing how code runs across multiple CPU cores.
// Type Checking: Verifies data types at runtime to prevent system crashes.
// System Interfacing: Mediates requests between your program and the operating system for tasks like reading files or printing text.
// How Go's Runtime Library DiffersUnlike languages like Java or Python which require a separate virtual machine (JVM) installed on the computer, 
// Go includes its runtime library directly inside the compiled binary executable file.Self-Contained: 
// Go compiles the runtime library straight into your final .exe or binary file.
// No Dependencies: The end user does not need to install Go to run your program.
// Low Overhead: Go's runtime library is exceptionally lightweight, managing rapid goroutine scheduling and efficient memory cleanup behind the scenes.