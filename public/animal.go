package public

import (
	"fmt"
)
	
type Animal struct {
	Name string
	Species string
	Age int
}

type Dog struct {
	Animal
	Race string
	Noise string
}

type Cat struct {
	Animal
	Race string
	Noise string
}

var Buddy = Dog{
	Animal: Animal{
		Name:    "Buddy",
		Species: "Canine",
		Age:     5,
	},
	Race:      "Labrador",
	Noise: "Woof",
}

var Test = Cat{
	Animal: Animal{
		Name:    "Test",
		Species: "Feline",
		Age:     2,
	},
	Race: "Persian", //private
	Noise: "meow",
}

func (d Dog) MakeANoise() {
	fmt.Println("Sound:", d.Noise)
}
func (c Cat) MakeANoise() {
	fmt.Println("Sound:", c.Noise)
}