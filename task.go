//given one number from a=3--->1,2,3
//a=3, b=2---->1,1,2,2,3,3
//a=5, b=3---->1,1,1,2,2,2,3,3,3,4,4,4,5,5,5
package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 3
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			fmt.Println(i + 1)
		}
		fmt.Println(" ")
	}
}
