package main

import "fmt"

func main() {
	//Widely used keyword

	//result := add(5, 6)
	//defer fmt.Println(result)
	add(2, 4)

	defer fmt.Println("This is 1st Line")
	fmt.Println("This is 2st Line")
	defer fmt.Println("This is 3st Line")
	fmt.Println("This is 4st Line")
	defer fmt.Println("This is 5st Line")

}

func add(a, b int) int {

	fmt.Println(a * b)
	return a * b
}
