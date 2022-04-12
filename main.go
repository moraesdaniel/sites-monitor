package main

import (
	"fmt"
	"os"
)

func main() {

	showMenu()

	menuOption := getChosenOption()

	if menuOption == 1 {
		fmt.Println("Starting monitoring...")
	} else if menuOption == 2 {
		fmt.Println("Showing logs...")
	} else if menuOption == 0 {
		fmt.Println("Exiting...")
		os.Exit(0)
	} else {
		fmt.Println("Invalid option!")
		os.Exit(-1)
	}
}

func showMenu() {
	version := 1.0

	fmt.Println("Welcome to the WebSites Monitor version", version)
	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Exit")
}

func getChosenOption() int {
	var menuOption int
	fmt.Scan(&menuOption) //The simbol "&" means that menuOption is a pointer of memory
	return menuOption
}
