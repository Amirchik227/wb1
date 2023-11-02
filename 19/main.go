package main

import "fmt"

func reverseString(s string) string {
	// Конвертируем строку в массив рун,
	// потому что напрямую со строкой работать не получится
	runes := []rune(s)

	// Получаем длину массива
	length := len(runes)

	// Переворачиваем массив
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-i-1] = runes[length-i-1], runes[i]
	}

	// Конвертируем в строку и возвращаем
	return string(runes)
}

func main() {
	s := "абырвалг"
	fmt.Println(s)

	// Выводим результат
	fmt.Println(reverseString(s))
}
