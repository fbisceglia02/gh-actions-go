The line `user := &User{name:"Frank"}` in Go does the following:

1. `User{name:"Frank"}`: This part of the line creates a new `User` struct with the `name` field set to `"Frank"`. It assumes that a struct type `User` has been defined somewhere in your code with a field named `name`.
   
2. `&`: This is the address-of operator, which gets the memory address of the struct created in step 1, effectively creating a pointer to the struct.

3. `user :=`: This part of the line declares a new variable named `user` and assigns it the pointer to the `User` struct created in steps 1 and 2.

So, to break it down:

- `User` is a struct type that must be defined elsewhere in your code.
- `name:"Frank"` initializes the `name` field of the `User` struct to the string "Frank".
- `&` creates a pointer to the `User` struct.
- `user :=` declares a new variable named `user` and assigns it the pointer to the `User` struct.

Here's how you might define the `User` struct and use the line in a complete Go program:

```go
package main

import "fmt"

type User struct {
    name string
}

func main() {
    user := &User{name:"Frank"}
    fmt.Println(user.name)  // Output: Frank
}
```

In this program:

- A new type `User` is defined as a struct with a single field named `name` of type `string`.
- In the `main` function, a new `User` struct is created with the `name` field set to "Frank", and a pointer to this struct is assigned to the variable `user`.
- `fmt.Println(user.name)` prints the `name` field of the `User` struct pointed to by `user`, which is "Frank".