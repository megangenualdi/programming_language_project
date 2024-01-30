1. What are the naming requirements for variables in your language?
  a. What about naming conventions? Are they enforced by the compiler/interpreter, or are they just standards in the community?

  Variables are camelcase and start with a lowercase letter unless they are being exported. These naming conventions in Go are not enforced by the complier but are best practices that are generally followed in the community. 
  
  Something that Go will enforce when it comes to variable names, is that they have to start with a letter, whether it's upper or lowercase. It will allow numbers in the variable name, they just cannot be at the first character. Special characters (with an exception of underscores) are not allowed in variable names.

  https://www.mohitkhare.com/blog/go-naming-conventions/

2. Is your language statically or dynamically typed?

  Go is statically typed because when a variable is first declared, it will continue to infer that variables type from the first assignment. So, for instance, you cannot assign an integer to a variable, then try to assign a string to that variable later on.

  https://stackoverflow.com/questions/73096744/how-is-golang-statically-typed-when-it-also-allows-not-specifying-any-type-for-a

3. Strongly typed or weakly typed?

  Go is strongly typed because it will only allow operations on certain variable types.

  https://www.techtarget.com/searchitoperations/tip/What-Golang-generics-support-means-for-code-structure#:~:text=For%20the%20most%20part%2C%20Go,implicitly%2C%20without%20any%20implements%20keyword.

4. If you put this line (or something similar) in a program and try to print x, what does it do? If it doesn't compile, why? Is there something you can do to make it compile? x = "5" + 6

  This is the error that appears when implementing the above statement in a go function: `invalid operation: "5" + 6 (mismatched types untyped string and untyped int)`. It will not compile because the `+` operator is being used on mismatched types and, as was said above, Go is a strongly typed language.
  Compiling solution:
    ```
    package main

    import (
	    "fmt"
	    "strconv"
    )

    func main() {
	    num,_ := strconv.Atoi("5")
	    var x = num + 6
	    fmt.Println(x)
    }

  https://www.techtarget.com/searchitoperations/tip/What-Golang-generics-support-means-for-code-structure#:~:text=For%20the%20most%20part%2C%20Go,implicitly%2C%20without%20any%20implements%20keyword.

5. Describe the limitations (or lack thereof) of your programming language as they relate to the coding portion of the assignment (adding ints and floats, storing different types in lists, converting between data types). Are there other restrictions or pitfalls that the documentation mentions that you need to be aware of?

  Coming from Python, I definitely miss working in a dynamically typed language. While I can see the benefit to only allowing methods/operators to work on/with the declared datatype, it feels like a limitation coming from a language that allows situations like reassigning datatypes and being able to store different datatypes in dicts/maps. Another thing that feels bulky compared to Python is string interpolation. Especially during debugging or creating descriptive error logs, it's beneficial to be able to quickly and simply write messages and not have to worry about the datatypes.

6. Are there built-in complex data types that are commonly used in your language? (hint: they’d probably appear fairly early in the documentation if so)

  There are a couple complex data types that I will need to be aware of: slices and structs. A slice is like an array, with the key difference being that an array's length needs to be defined, while a slice can remain undefined.

  Structs are like maps, where they store key value pairs but they can be used to store more complex data and data that is of different datatypes but still related.

  https://medium.com/@diwakarkashyap/complex-data-types-in-golang-go-6410f136bad7

