package main

import "fmt"

func addition(num1 int, num2 int) int {
	var result int

	result = num1 + num2

	return result
}
func main() {
	num1 := 12
	num2 := 12
	var result int

	result = addition(num1, num2)

	fmt.Println("Addition is: ", result)
}
