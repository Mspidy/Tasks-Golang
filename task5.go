//Create a program  to print "Hello <username>" from an array of usernames

package main

import (
	"fmt"
)

func main() {
	// var i int
	var username = [6]string{"Prabhat", "Avi", "Anand", "Adi", "Raj", "Tilak"}

	for i := 0; i < 5; i++ {
		fmt.Println("Hello", username[i])
	}
}
