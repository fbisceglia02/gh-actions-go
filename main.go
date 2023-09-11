package main

import "fmt"
import "os"
// import "public/person.go"

// Define a new type named Person




func main() {
	fmt.Println("hello from go")

	// p := person.Person{name:"Alice",age:20,isAdult:true}

	// fmt.Println("by ",p.name)

	directory, err := os.Getwd()    //get the current directory using the built-in function
	if err != nil {
	   fmt.Println(err) //print the error if obtained
	}
	fmt.Println("Current working directory:", directory) //print the required directory	
}


