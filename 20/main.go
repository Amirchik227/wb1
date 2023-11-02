package main

import "fmt"

func revereseWords(s string) string {
	// В данной карте будут хрнится слова из входной строки
	// Ключом каждой строки будет ее порядковый номер
	wordsMap := make(map[int]string)
	// Поиск слов будет происходить по символу пробела
	// Чтобы не добавлять отдельной логики для обработки последнего слова добавим пробел
	s += " "
	// Количество слов
	var wordsCount int
	// Индекс первого байта слова
	var start int
	for idx, val := range s {
		if val == 32 { // 32 - ascii код пробела
			wordsCount++
			wordsMap[wordsCount] = s[start:idx]
			start = idx + 1
		}
	}

	var reversed string
	// Достаём из мапы слова в отбратном порядке и добавляем
	// к результирующей строке через пробел
	for wordsCount > 1 {
		reversed += wordsMap[wordsCount] + " "
		wordsCount--
	}
	// После крайнего слова пробел не нужен
	if wordsCount == 1 {
		reversed += wordsMap[wordsCount]
	}

	return reversed
}

func main() {
	str := "My name is アミール"
	fmt.Println(str)
	fmt.Println(revereseWords(str))
}
