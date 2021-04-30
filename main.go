package main

import "fmt"

func main() {
	fmt.Println("Best Program....")
	fmt.Println("New feature!")
	n := 10
	fmt.Printf("%v is a square", square(n))
}

func square(number int) int {
	return number * number
}
