package public

type Person struct {
    Name string
    Age  int
	IsAdult bool
}

var John = Person{
    Name: "John",
    Age: 20,
    IsAdult: true,
}