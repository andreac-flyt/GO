package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// readName()
	// readUser()
	// readInput()
	// readMore()
	scanInput()
}

func readName() {

	var name string

	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Println("Hello,", name)

}

func readUser() {

	var name string
	var age int
	var city string

	fmt.Print("What is your name?: ")
	fmt.Scanf("%s", &name)
	fmt.Print("How old are you,", name, "? ")
	fmt.Scanf("%d", &age)
	fmt.Print("What city do you live in?: ")
	fmt.Scanf("%s", &city)
	fmt.Printf("Your name is %v, you are %v years old, and you live in %v.\n", name, age, city)

}

func readInput() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("What is your name?: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s\n", name)

}

func readMore() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Tell me your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s\n", name)

}

func scanInput() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your city?: ")
	scanner.Scan()
	location := scanner.Text()
	fmt.Printf("You live in %v\n", location)
}
