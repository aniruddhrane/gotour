A function can take two or more arguments
(in go the type comes after the variable name)
when there are consecutive parameters of the same type then we mention the type of all parameters int the last argument 

func add(x, y int) int {
	return x + y
}

named return values 
Go's return values may nbe named.
They are treated as variable defined at the top of the function

A return without arguments returns the named return values.
This is known as naked return.