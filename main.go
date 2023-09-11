package main

import "fmt"

// Define a new type named Person
type Person struct {
    Name string
    Age  int
}

// Define a method named Describe on the Person type
func (p Person) Describe() {
    fmt.Printf("My name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    // Create a new Person object
    p := Person{Name: "Alice", Age: 20}

    // Call the Describe method on the Person object
    p.Describe()

	// slice are dynamic arrays in go
	// a slice of persons:
	var persons []Person

	// append element to the array
	persons = append(persons, p)
}


