package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Используем библиотеку math/big для работы с большими числами
	// С её помощью создадим новый объект типа big.Int. Он будет
	// использоваться для хранения большого целого числа,
	// значение которого > 2^20
	a := new(big.Int)
	// Задаем значение а, как большое целое число из строки в 10-ой системе счисления
	a.SetString("98765432198765432198765", 10)

	// Аналогично для b
	b := new(big.Int)
	b.SetString("12345678912345678912345", 10)

	// Сложение больших чисел
	addition := new(big.Int)
	// С помощью метода Add сложим числа и сохраним в переменной addition
	addition.Add(a, b)
	fmt.Println("Сложение:", addition)

	// Вычитание больших чисел
	subtraction := new(big.Int)
	subtraction.Sub(a, b)
	fmt.Println("Вычитание:", subtraction)

	// Умножение больших чисел
	multiplication := new(big.Int)
	multiplication.Mul(a, b)
	fmt.Println("Умножение:", multiplication)

	// Деление больших чисел
	division := new(big.Int)
	division.Div(a, b)
	fmt.Println("Деление:", division)
}
