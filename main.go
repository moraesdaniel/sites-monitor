package main

import "fmt"

func main() {
	version := 1.0

	fmt.Println("Welcome to the WebSites Monitor version", version)
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")

	var menuOption int

	fmt.Scan(&menuOption) //The simbol "&" means that menuOption is a pointer of memory
	fmt.Println("The menu you chose is", menuOption)

	if menuOption == 1 {
		fmt.Println("Starting monitoring...")
	} else if menuOption == 2 {
		fmt.Println("Showing logs...")
	} else if menuOption == 0 {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Invalid option!")
	}
}
