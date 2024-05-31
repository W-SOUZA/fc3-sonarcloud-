package main

import "fmt"

func main() {
	fmt.Println(add(2, 2))
	fmt.Println(subtract(5, 3))
	fmt.Println(multiply(3, 4))
	fmt.Println(divide(10, 2))
	fmt.Println(modulo(10, 3))
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func modulo(a, b int) int {
	return a % b
}
