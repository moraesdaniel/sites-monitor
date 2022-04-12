package main

import "fmt"

func main() {
	version := 1.0

	fmt.Println("Welcome to the WebSites Monitor version", version)
	fmt.Println("1 - Initialize monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("3 - Exit")

	var menuOption int

	fmt.Scan(&menuOption) //The simbol "&" means that menuOption is a pointer of memory
	fmt.Println("The menu you chose is", menuOption)
}
