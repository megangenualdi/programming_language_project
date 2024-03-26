### 1. Does your language support objects or something similar (e.g., structs)? Are there naming conventions for objects, instance variables, or functions that people writing in your language should be aware of?

Go supports structs. The naming conventions for structs are that they begin with a capital letter and are camelcase. The instance variables can begin with an upper or lowercase letter, the same goes for functions. If you want to call your struct variables or functions outside of the file they are declared in, though, then they have to begin with a capital letter inorder to be imported.

If some of your struct's variables are of the same type, you can condense them to be on one line and comma separated, otherwise each variable will be on a new line.

```
type Person struct {
	Name, favoriteFood string
	Age int
}
```

In the above example, I can use the `Person` struct outside of this file because if is capatilized. I can also call the attributes `Name` and `Age` for the same reason. I would not be able to access the attribute `favoriteFood`, though, as it begins with a lowercase letter.

Another thing to be aware of is that if a struct method alters an attribute of the struct, in order for the alteration to persist, when the struct the method is for is declared, it has to be percided by a `*` to point to class.

```
// Will run without errors but won't save the new name
func (p Person) SetName(newName string) {
	p.Name = newName
}

// Will actually save the new name
func (p *Person) SetName(newName string) {
	p.Name = newName
}
```
***
### 2. Does your language have standard methods for functions that serve a similar purpose across all objects? For example, toString() in Java and __str__ in Python allow information about the objects to be printed. Are there similar functions in your language?

There are similiar functions to Go. To achieve the same results as `toString()` or `__str__` you would do the below:

```
func (p *Person) String() string {
 return fmt.Printf("Name: %s, Age: %d, Favorite food: %s",
                    p.Name, p.Age, p.FavoriteFood)
}
```

***
### 3. How does inheritance work (if it does)? Does your language support multiple inheritance?

Classes are not supported in Go, so inheritance, or it's Go equivalent, is achieved through struct embedding. When struct is used(embedded) within another struct, the embedded structs functions and attributes are not readily available to the other struct. In order to use the embedded struct's attributes/functions you have to 'chain' call the embedded(parent) struct from your 'child' struct.

Example:
```
type Employee struct {
	Person
	JobTitle string
}

employee := classes.Employee{
	Person: classes.Person{Name: "Worker B"},
}
employee.Person.SetAge(24)
employee.SetJob("Software Engineer")
```
***
### 4. If there is inheritance, how does your language deal with overloading method names and resolving those calls?

Because Golang uses embedding for it's inheritance, methods are not overloaded.

```
type Person struct {
	Name, FavoriteFood string
	Age int
}

func (p Person) Introduction() {
	fmt.Printf("My name is %s, my favorite food is %s and my age is %d", p.Name, p.FavoriteFood, p.Age)
}

type Employee struct {
	Person
	JobTitle string
}

func (e Employee) Introduction() {
	// e.Person.Introduction()
	fmt.Printf("\nMy name is %s, my job title is %s\n", e.Name, e.JobTitle)
}
```

***
### 5. Is there anything else that’s important to know about objects and inheritance in your language?

Not necessarily having to do with objects and inheritance but it is a bit different to import code from one file to another and I feel like it's work talking about:

Step One: 
- Create your go.mod file `go mod init <name of directory>`

Step Two:
- At the top of the file that contains the code you want to import declare the package name for this file: `package <package name>`

Step Three:
- In the file that you want to import too, have an import that looks like this:
```
import (
	"<name of directory>/<package name>"
)
```

References

- https://www.geeksforgeeks.org/inheritance-in-golang/

- https://www.geeksforgeeks.org/structures-in-golang/

- https://stackoverflow.com/questions/47357201/how-to-modify-struct-fields-in-golang

- https://stackoverflow.com/questions/37780520/unknown-field-in-struct-literal

- https://stackoverflow.com/questions/34644117/golang-struct-inheritance-not-working-as-intended

- https://www.tutorialspoint.com/golang-program-to-show-inheritance-in-class#:~:text=As%20classes%20are%20not%20supported,Golang%20doesn't%20support%20inheritance.

- https://fullstackdojo.medium.com/the-python-str-method-and-the-go-string-method-a-comparison-75770c78c4d8
