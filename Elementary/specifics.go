package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	// alice_bob()
	dayGreeting()
}

func alice_bob() {

	var name string

	fmt.Print("Type your first name: ")
	fmt.Scanf("%s", &name)
	testName := strings.ToLower(name)
	if testName == "alice" {
		fmt.Printf("Hello, %s\n", testName)
	} else if testName == "bob" {
		fmt.Printf("Hello, %s\n", testName)
	} else {
		fmt.Print("Goodbye.\n")
	}

}

func dayGreeting() {

	var first string
	var last string

	fmt.Print("Please state your first name: ")
	fmt.Scanf("%s", &first)
	fmt.Print("Please state your last name: ")
	fmt.Scanf("%s", &last, '\n')
	fullName := fmt.Sprintf("%s %s", first, last)

	weekday := time.Now().Weekday()

	fmt.Printf("Hello %s, I hope you have had a great %s.\n", fullName, weekday)

}
