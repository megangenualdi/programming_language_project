// Write a function that takes in a string (or your language's equivalent) and splits it
// into two strings, then returns both strings
package main

import (
	"fmt"
	"strings"
)


func main() {
	strOne, strTwo := splitStrs("Split this")
	fmt.Println(strOne)
	fmt.Println(strTwo)
}


func splitStrs(str string) (string, string) {
	split := strings.Split(str, " ")
	return split[0], split[1]
}

// https://www.geeksforgeeks.org/how-to-split-a-string-in-golang/
