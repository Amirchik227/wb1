package main

import "fmt"

// Удаление i-ого элемента из слайса
func deleteElement(s []int, idx int) []int {
	// Удаляем элемент с заданным индексом
	// Выберем чать массива до i-го слайда и добавим в неё все значения после (i+1)-го элемента,
	// которые мы распаковываем оператором ...
	s = append(s[:idx], s[idx+1:]...)
	return s
}

func main() {
	// Инициализация слайса
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr)

	// Слайс после удаления i-го элемента
	arr = deleteElement(arr, 5)
	fmt.Println(arr)
}
