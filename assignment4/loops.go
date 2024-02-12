package main

import "fmt"


func main() {
	forLoop := forLoopExample("backwards")
	fmt.Println(forLoop)

	whileLoop := whileLoopExample(56)
	fmt.Println(whileLoop)

	rangeLoop := rangeLoopExample()
	fmt.Println(rangeLoop)
}


func forLoopExample(str string) string {
	reversedStr := ""
	for i:=0; i<len(str); i++{
		reversedStr = string(str[i]) + reversedStr
	}
	return reversedStr
}


func whileLoopExample(num int) int {
	dividedNum := num 
	for dividedNum > 2 {
		dividedNum = dividedNum / 2
	}
	return dividedNum
}


func rangeLoopExample() bool {
	str := "abcdefghijklmnopqrstuvwxyz"

	for i, x := range str {
		fmt.Println(i, string(x))
	}
	return true
}


func infiniteLoopExample() {
	for {
		fmt.Println("This will go forever until it's stopped manually")
	}
}

// https://www.codecademy.com/resources/docs/go/loops
