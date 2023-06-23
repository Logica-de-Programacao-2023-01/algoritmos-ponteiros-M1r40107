package main

import "fmt"

func updateIntValue(ptr *int) {
	num := *ptr
	UltimoDigito := num % 10
	num /= 10
	SegundoUltimoDigito := num % 10
	sum := UltimoDigito + SegundoUltimoDigito
	*ptr = sum
}

func main() {
	num := 1234
	updateIntValue(&num)
	fmt.Println("Updated value:", num)
}
