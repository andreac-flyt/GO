package main

import "fmt"

func main() {

	addNums()
	addAll()
}

func addNums() {

	var num int

	fmt.Print("Enter a whole number: ")
	fmt.Scanf("%d", &num)
	sum := 0
	for i := 1; i < num; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func addAll() {

	var n int
	total := 0

	fmt.Print("Please enter a whole number: ")
	fmt.Scanf("%v", &n)
	for i := 1; i < n; i++ {
		total += i
	}
	fmt.Printf("The total sum of all numbers from 1 to %v is %v.\n", n, total)
}
