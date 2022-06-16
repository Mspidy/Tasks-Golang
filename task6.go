//Create a program  to print "Hello my name is  '<username>'. I am a student".  From an array of students.
//Notice the single quote surrounding the <username>.
package main

import (
	"fmt"
)

func main() {
	// var i int
	var username = [6]string{"Prabhat", "Avi", "Anand", "Adi", "Raj", "Tilak"}

	for i := 0; i < 5; i++ {
		fmt.Println("Hello my name is " + username[i] + ". I am a student")
	}
}
