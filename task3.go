//Create a program  to print "Hello <username>" from an array of usernames
package main

import "fmt"

func main() {
	var username string
	//taking input from user
	fmt.Println("Enter username:")
	fmt.Scanln(&username)

	fmt.Print("hello ", username)
}
