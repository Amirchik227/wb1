package main

import (
	"fmt"
	"strings"
)

func isUnique(s string) bool {
	// Переводим строку в нижний регистр
	s = strings.ToLower(s)
	// Создаем карту для проверки уникальности символов в строке
	charMap := make(map[rune]bool)

	for _, v := range s {
		// Если символ уже с встречался, то не все символы в строке уникальны
		if charMap[v] { 
			return false
		}
		// Записываем в карту что данный символ присутствует в строке
		charMap[v] = true
	}

	// Если при переборе символов не нашлось повторов, значит все они уникальны
	return true
}

func main() {
	// Инициализируем строку
	str := "Unique"

	// Выводим результат
	fmt.Println(isUnique(str))
}
