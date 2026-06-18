the go program is made up of packages
the go program starts running in package main:
the convention is to keep the package name same as the last element of the import path 

suppose the "math/rand" package comprises files that begin with the statement package rand.

in go files are just container they dont have independent scope like modules has in node js

package: all.go files in a single folder sharing the same package name.
all the files in that same package(directory) share same namespace 

a package is a unit of code organization not file
files are merely containers 
the package determines the namespace

import "math/rand" is the import path.
Inside that folder/package, the actual package name written in the Go files is usually the last part of the path, so for math/rand the package name is rand.

So this line:

import "math/rand"

lets you use that package as:

rand.Intn(10)

because the package name is rand, not math/rand.

In your program
package main

import (
	"fmt"
	"math/rand"
)
fmt is a package, and its package name is also fmt
math/rand is the import path, but the package name is rand

So:

fmt.Println(...)

uses the fmt package.

rand.Intn(10)

uses the rand package and calls its Intn function.