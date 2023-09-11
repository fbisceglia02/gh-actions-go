package main

import "fmt"
import "public/person.go"

// Define a new type named Person




func main() {
	fmt.Println("hello from go")

	p := person.Person{name:"Alice",age:20,isAdult:true}

	fmt.Println("by ",p.name)
}


