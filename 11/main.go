package main

import "fmt"

func main() {
	// Создаем два массива
	var array1 = [...]int{1, 43, 6, 255, 23, 991, 235, 22, 6, 2100, 170}
	var array2 = [...]int{1, 43, 7, 255, 23, 100, 769, 22, 0}

	// Слайс для пересечения множеств
	var intersection []int

	// Карта для построения пересечения
	var countMap = make(map[int]int)

	// Возьмем все элементы первого массива в качестве ключей для карты и
	// положим по ним единицы.
	for _, value := range array1 {
		countMap[value] = 1
	}
	// Возьмем все элементы второго массива в качестве ключей для карты и,
	// если по данным ключам уже лежали единицы (таким образом ключ являтся
	// элементом и первого и второго массива), положим по ним двойки.
	for _, value := range array2 {
		if countMap[value] > 0 {
			countMap[value] = 2
		}
	}

	// Все элементы, входяие в оба массива добавим в результирующий слайс
	for key, value := range countMap {
		if value == 2 {
			intersection = append(intersection, key)
			fmt.Println(key)
		}
	}
}
