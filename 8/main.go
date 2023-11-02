package main

import "fmt"

func setBit(number int64, index int, value int) int64 {
	// Операция 1 << index позволяет создать число в двоичной записи которого 
	// от нуля отличен лишь бит по индексу index
	if value == 1 {
		// Логическое ИЛИ обеспечит в нужном инексе единицу
		number |= 1 << index
	} else {
		// Оператор НЕ (^) инвертирует все биты исходного числа
		// Применив после логическое И обеспечим 0 в необходимом нам бите 
		number &= ^(1 << index)
	}
	return number
}

func main() {
	fmt.Println(setBit(16, 2, 1))
	fmt.Println(setBit(16+8, 3, 0))
}
