package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const qtyMonitoring = 3
const secondsToSleep = 10

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
	fmt.Println("##########################################")
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
	fmt.Println("Starting monitoring " + strconv.Itoa(qtyMonitoring) + "x...")

	websites := getSitesFromFile("./websites.txt")

	for loop := 0; loop < qtyMonitoring; loop++ {
		for _, site := range websites {
			testWebsite(site)
		}
		time.Sleep(secondsToSleep * time.Second)
	}
}

func testWebsite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if response.StatusCode == 200 {
		fmt.Println("Website", site, "loaded successfully!")
	} else {
		fmt.Println("Website", site, "has a problem! Status code:", response.StatusCode)
	}
}

func getSitesFromFile(fileName string) []string {
	var websites []string

	fileContent, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error:", err)
	}

	reader := bufio.NewReader(fileContent)

	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)

		if err != nil && err != io.EOF {
			fmt.Println("Error:", err)
		}

		websites = append(websites, site)

		if err == io.EOF {
			break
		}
	}

	fileContent.Close()

	return websites
}
