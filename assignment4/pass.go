// Write a function that tests whether your language is pass-by-reference or pass-by-value.
package main

import "fmt"


type Person struct {
	Name string
	FavoriteFood string
	Age int
}

func (p Person) Introduction() {
	fmt.Printf("My name is %s, my favorite food is %s and my age is %T", p.Name, p.FavoriteFood, p.Age)
}


func main() {
	person := Person{
		Name: "Bob",
		FavoriteFood: "Pizza",
		Age: 33,
	}
	passByValue(person)
	fmt.Println("\nAge and name called within main function")
	fmt.Printf("%s %d\n======>\n", person.Name, person.Age)

	passByReference(&person)
	fmt.Println("\nAge and name called within main function")
	fmt.Printf("%s %d\n======>\n", person.Name, person.Age)
}

func passByValue(person Person) {
	person.Age += 10
	person.Name = "Bobert"
	fmt.Printf("This is Bob's age and name called within passByValue: %s %d",
person.Name, person.Age)
}

func passByReference(person *Person) {
	person.Age += 10
	person.Name = "Bobert"
	fmt.Printf("This is Bob's age and name called within passByReference: %s %d",
person.Name, person.Age)
}

// This is Bob's age and name called within passByValue: Bobert 43
// Age and name called within main function
// Bob 33
// ======>
// This is Bob's age and name called within passByReference: Bobert 43
// Age and name called within main function
// Bobert 43
// ======>
