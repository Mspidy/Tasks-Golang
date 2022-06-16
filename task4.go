//Create an array of 10 integers. Pass the array as input to a function called "AddArray"
//which accepts as input an array and
//it's size and returns the sum of all elements in the array

package main

import "fmt"

func myfun(a [6]int, size int) int {
	var k, val, r int

	for k = 0; k < size; k++ {
		val += a[k]
	}

	r = val / size
	return r
}

func main() {
	var arr = [6]int{67, 59, 29, 35, 4, 34}
	var res int

	res = myfun(arr, 6)
	fmt.Printf("Final result is: %d ", res)
}
