package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func resta(a, b int) int {
	return a - b
}

func main() {
	fmt.Println(suma(6, 3))
	fmt.Println(resta(8, 8))
}
