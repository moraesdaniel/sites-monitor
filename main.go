package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		showMenu()

		menuOption := getChosenOption()

		if menuOption == 1 {
			startMonitoring()
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

func startMonitoring() {
	fmt.Println("Starting monitoring...")
	site := "https://random-status-code.herokuapp.com/"
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Website", site, "loaded successfully!")
	} else {
		fmt.Println("Website", site, "has a problem! Status code:", response.StatusCode)
	}
}
