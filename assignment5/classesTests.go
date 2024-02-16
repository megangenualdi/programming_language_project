package main

import (
	"fmt"
	"assignment5/classes"
)

func main() {
	person := classes.Person{
		Name: "Jane",
		Age: 33,
	}
	fmt.Println(person) // {Jane  33} 
	person.SetName("Jane Person")
	person.SetAge(45)
	person.SetFavoriteFood("Pizza")
	fmt.Println(person) // {Jane Person Pizza 45}

	employee := classes.Employee{
		Person: classes.Person{Name: "Worker B"},
	}
	fmt.Println(employee) // {{Worker B  0} }

	employee.Person.SetAge(24)
	employee.SetJob("Software Engineer")
	fmt.Println(employee) // {{Worker B  24} Software Engineer}

	employee.Introduction() // My name is Worker B, my job title is Software Engineer
	employee.Person.Introduction() // My name is Worker B, my favorite food is  and my age is 24
}

