package main

import "fmt"

func main() {
	Search([6]int{67, 59, 29, 35, 4, 34}, 35)
}

func Search(a [6]int, size int) {
	for i := 0; i < len(a); i++ {
		if a[i] == size {
			fmt.Println(i)
		} else {
			fmt.Println(-1)
		}
	}
}
