Go Pointers, Stack, Heap and Goroutines
What is a Pointer?

A pointer is a variable that stores the memory address of another variable.

Instead of storing the actual value, it stores the location where the value is stored.

x := 42
p := &x
x = 42      Address: 0x100

p = 0x100

Here:

&x → address of x
p → pointer storing that address
*p → value at that address
fmt.Println(*p) // 42

*p = 100

fmt.Println(x) // 100

Because p points directly to x, modifying *p modifies x.

Goroutines

A goroutine is a lightweight unit of execution managed by the Go runtime.

Features:

Cheaper than OS threads
Thousands of goroutines can run concurrently
Created using the go keyword
Each goroutine has its own stack

Example:

go myFunction()
Stack

When a function is called, Go creates a stack frame.

A stack frame contains:

Function parameters
Local variables
Return information

Example:

func square(v int) {
    v = v * v
}

a := 4
square(a)

Memory:

main frame
-----------
a = 4

square frame
------------
v = 4

v receives a copy of a.

Therefore:

fmt.Println(a) // 4

The original variable is unchanged.

Important

Go always passes arguments by value.

That means a copy is passed to the function.

Why Do We Need Pointers?

Without pointers:

func update(n int) {
    n = 100
}

x := 10
update(x)

fmt.Println(x) // 10

Only the copy changes.

With pointers:

func update(n *int) {
    *n = 100
}

x := 10
update(&x)

fmt.Println(x) // 100

Now the original variable changes.

Main Uses of Pointers
1. Modify Original Data
func update(n *int)

Allows modification of the caller's variable.

2. Avoid Copying Large Structs
type User struct {
    Name string
    Age  int
}

Instead of:

func process(u User)

use:

func process(u *User)

Only an address is copied.

3. Share Data Between Functions

Multiple functions can access the same object.

4. Mutating Methods
func (u *User) UpdateAge(age int)

Changes the actual object.

Understanding *

The * symbol has two meanings.

1. Pointer Type
var p *int

*int means:

Pointer to an integer

Here *int is a type.

2. Dereference Operator
*p

Means:

Give me the value stored at the address inside p

Example:

x := 42
p := &x

fmt.Println(*p) // 42
Pointer Example
i, j := 42, 2701

fmt.Println(&i, &j)

p := &i

fmt.Println(*p)

p = &j

*p = *p / 37

fmt.Println(j)

Flow:

p -> i -> 42

then

p -> j -> 2701

*p = *p / 37

2701 / 37 = 73

j becomes 73
Trade-Offs of Pointers
Advantages

✅ Less memory copying

✅ Better performance for large structs

✅ Can modify original data

✅ Shared access to data

Disadvantages

❌ Harder to reason about code

❌ Functions can unexpectedly modify data

❌ More bugs from shared mutable state

❌ May cause heap allocations

❌ More work for the Garbage Collector
Heap

The heap is a region of memory used to store data that must survive beyond the lifetime of a function call.

Unlike stack memory, heap memory is not automatically removed when a function returns.

Why Do We Need the Heap?

Consider:

func create() *int {
    x := 10
    return &x
}

At first glance:

create() frame
--------------
x = 10

When the function returns, its stack frame disappears.

If x remained on the stack:

return &x

would return an address pointing to invalid memory.

To prevent this, Go moves x to the heap.

Escape Analysis

Go uses escape analysis to determine whether a variable should be allocated on the stack or heap.

Example:

func create() *int {
    x := 10
    return &x
}

Since x is needed after the function returns:

x escapes the stack

Therefore:

Heap
----
x = 10

Stack
-----
pointer -> x

The returned pointer remains valid.

Stack vs Heap
Stack	Heap
Fast allocation	Slower allocation
Automatic cleanup	Garbage Collector cleanup
Per goroutine	Shared process memory
Function-local data	Long-lived data
No GC overhead	GC overhead
Stack Example
func square(v int) int {
    return v * v
}
main frame
-----------
a = 4

square frame
------------
v = 4

After square() returns:

square frame removed

Memory is automatically reclaimed.

Heap Example
type User struct {
    Name string
}

func NewUser() *User {
    u := User{Name: "Annie"}
    return &u
}

Since u must exist after NewUser() returns:

Heap
----
User{Name:"Annie"}

Stack
-----
pointer -> User

Go automatically places u on the heap.

Garbage Collector (GC)

Heap memory is cleaned by Go's Garbage Collector.

The GC periodically:

Finds objects still being referenced.
Keeps those objects.
Frees unreachable objects.

Example:

u := NewUser()

u = nil

Now nothing points to the User object.

The GC will eventually remove it from memory.

Heap and Pointers

A common misconception:

Pointer => Heap
Value => Stack

❌ Wrong

Example:

x := 10
p := &x

Both can remain on the stack.

Stack
-----
x = 10
p -> x

No heap allocation is required.

When Does Heap Allocation Happen?

Heap allocation occurs when data must outlive its current stack frame.

Common cases:

return &x
return &User{}
store pointer in a long-lived object
capture variable inside a closure

The compiler decides this using escape analysis.

Heap and Pointers Trade-off

Using pointers may:

✅ Reduce copying

✅ Improve performance for large structs

❌ Cause heap allocations

❌ Increase GC work

❌ Increase memory pressure

Therefore:

Use pointers when you need shared mutable data or want to avoid expensive copies—not simply because "pointers are faster."