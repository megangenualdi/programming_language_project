// Write a recursive function (one that calculates a factorial is fine)
package main

import "fmt"


func main() {
	num := recursiveFunc(56)
	fmt.Println(num)
}


func recursiveFunc(num int) int {
	if num > 2 {
		num = recursiveFunc(num / 2)
	}
	return num
}
