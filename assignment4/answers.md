### 1. Does your language include multiple types of loops (while, do/while, for, foreach)? If so, what are they and how do they differ from each other?

Go does support multiple types of loops (`for` and `while`) but they are both implemented by using the `for` keyword.
***

### 2. What is the syntax for declaring a function in your language?

A function is declared with the `func` keyword, followed by the function name. Then in parenthesises after the name is where arguments are declared, if any. An argument name is required, as well as the argument type. Afterwards, if the function is returning something, the return variable's type needs to be declared. Functions are enclosed using `{ }`.

Example
```
func recursiveFunc(num int) int {
	if num > 2 {
		num = recursiveFunc(num / 2)
	}
	return num
}
```
***

### 3. Are there any rules about where the function has to be placed in your code file so that it can run?

No, a function can be placed before or after where it is being called and it will still run. There are also no rules about a main function having to be the first one declared in a file either.
***

### 4. Does your language support recursive functions?

Yes, recursive functions are supported.
***

### 5. Can functions in your language accept multiple parameters? Can they be of different data types?

Yes functions can except multiple parameters and they can be different data types.
***

### 6. Can functions in your language return multiple values at the same time? How is that implemented? If not, are there ways around that problem? What are they?

Yes functions can return multiple values at the same time. These functions are implemented the same way as stated above, with the difference being when you are declaring the return value type, you use parenthesises and provide the comma separated values for each value type.

Example
```
func splitStrs(str string) (string, string) {
	split := strings.Split(str, " ")
	return split[0], split[1]
}
```
***

### 7. Is your language pass-by reference or value? Check your code against outside sources in case there is anything tricky going on (like in Perl).

Golang is able to support pass-by reference and pass-by value. I have gone more in depth with this in the next question.
***

### 8. Are there any other aspects of functions in your language that aren't specifically asked about here, but that are important to know in order to write one? What are they?

One important aspect has to do with declaring function arguments and their variable types. There are two ways to declare argument types and they will perform differently as a result.

Option 1 (Pass by value)
```
func randomFunction(argOne Random) {
	return "This is a function who's argument type is 'Random'"
}
```
The above function'sa argument means that argOne's variable type is the pretend class Random

Option2 (Pass by reference)
```
func randomFunction(argOne *Random) {
	return "This is a function who's argument type is POINTED at 'Random'"
}
```
When are argument type has an asterik in front of it, it means that the argument is pointed at that class type (in our example this is Random) and if argOne is manipulated in any way within our randomFunction then those changes will persist outside of the function. Also, when randomFunction is invoked and is pass by reference the argument being passed to it needs to be precided by an `&`
Example
```
randomFunction(&argOne)
```

Another characteristic to be mindful of is that function argmuents cannot be set to a default value like in other languages, like Python. There are workarounds for this, like immediately checking the argument to see if it is null and then assigning it to your default value.
Example
```
func someFunction(needsDefault string) {
   if needsDefault == "" {
       needsDefault = "This is my default value"
   }
}
```
***

References:

- https://www.codecademy.com/resources/docs/go/loops

- https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions#:~:text=First%2C%20Go%20technically%20has%20only,is%20subtle%20but%20occasionally%20relevant.

- https://www.geeksforgeeks.org/how-to-split-a-string-in-golang/

- https://david-yappeter.medium.com/golang-pass-by-value-vs-pass-by-reference-e48aac8b2716

- https://golangbot.com/structs-instead-of-classes/

- https://stackoverflow.com/questions/19612449/default-value-in-gos-method
