// • int
// • string
// • floating-point number
// • boolean
// • array/list
// • dictionary (sometimes called a hash or a map, depending on your language)
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "This is a string"
	var num = 34
	var floatNum = 100.995
	var boolVar = true
	arrList := [...]string{"Apples", "Bananas", "Potato"}
	mapObj := map[string]string{
		"name": "Megan",
		"language": "Go",
		"mood": "hungry",
	}
	// Arrays and Maps can contain only one datatype
	fmt.Printf("%s\n%d\n%f\n", str, num, floatNum)
	// 	This is a string
	// 34
	// 100.995000true
	fmt.Println(boolVar)
	// true
	fmt.Println(arrList)
	// [Apples Bananas Potato]
	fmt.Println(mapObj)
	// map[language:Go mood:hungry name:Megan]

	// sum := num * floatNum throws a mismatched types error, one solution: convert float to int
	sum := num * int(floatNum)
	fmt.Printf("sum = %d and is type %T", sum, sum)
	// sum = 3400 and is type int

	newNum := strconv.Itoa(num)
	// cannot reassign `num` as a string, so I had to create a new variable
	fmt.Printf("newNum = %s and is type %T", newNum, newNum)
	// newNum = 34 and is type string

	backToInt,_ := strconv.Atoi(newNum)
	fmt.Printf("backToInt = %d and is type %T", backToInt, backToInt)
	// backToInt = 34 and is type int
}
