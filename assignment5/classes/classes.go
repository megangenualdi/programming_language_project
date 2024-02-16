package classes

import "fmt"


type Person struct {
	Name, FavoriteFood string
	Age int
}

func (p *Person) SetName(newName string) {
	p.Name = newName
}

func (p *Person) SetFavoriteFood(food string) {
	p.FavoriteFood = food
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) Introduction() {
	fmt.Printf("My name is %s, my favorite food is %s and my age is %d", p.Name, p.FavoriteFood, p.Age)
}


type Employee struct {
	Person
	JobTitle string
}

func (e *Employee) SetJob(job string) {
	e.JobTitle = job
}

func (e Employee) Introduction() {
	// e.Person.Introduction()
	fmt.Printf("\nMy name is %s, my job title is %s\n", e.Name, e.JobTitle)
}
