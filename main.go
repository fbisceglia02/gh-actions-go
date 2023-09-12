package main

import "fmt"
import "go-lang/public"

// Define a new type named Person


// func printGreeting() {
// 	fmt.Println("hello from go")

// 	p := person.Person{name:"Alice",age:20,isAdult:true}

// 	fmt.Println("by ",p.name)
// }

func main() {
	// fmt.Println(public.Buddy.race
	public.Buddy.MakeANoise()
	public.Test.MakeANoise()

	fmt.Println(public.Test.Name,"is a",public.Test.Race,public.Test.Species)
}

