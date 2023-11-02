package main

import "fmt"

func main() {
	// Создаем два числа
	a, b := 22, 79
	fmt.Printf("a=%d, b=%d\n", a, b)

	// Меняем их значения местами
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)
	// Меняем их значения местами еще раз
	a ^= b
	b ^= a
	a ^= b
	fmt.Printf("a=%d, b=%d\n", a, b)

	// Меняем их значения местами еще раз
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("a=%d, b=%d\n", a, b)
}
