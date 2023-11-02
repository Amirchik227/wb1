package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Комментраии в решении 16-й задачи
func generateSlice(size int) []int {
	slice := make([]int, size)
	rg := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < size; i++ {
		slice[i] = rg.Intn(300)
	}
	return slice
}

// Комментраии в решении 16-й задачи
func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1
	pivot := rand.Int() % len(a)
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	quicksort(a[:left])
	quicksort(a[left+1:])
	return a
}

func binarySerach(arr []int, s int) int {
	// Определяем левую и правую границы массива
	left, right := 0, len(arr)-1
	// Далее будем сужать подмассив поиска зачения пока
	// индекс левой границы <= индекса правой
	for left <= right {
		// Возьмем элемент примерно в середине массива
		mid := int((left + right) / 2)
		// Найдем его значение
		val := arr[mid]
		// Если значение совпадает с тем, индекс которого мы ищем, то задача решена
		if val == s {
			return mid
			// Если значение меньше нужного, можно не рассматривать левую часть
			// в текущем массиве (отсортированном), так все его значения не подходят,
			// ведь все они меньше нужного (как и значеие по индексу mid). Далее
			// свдвинем левую границу до значния следующего за mid и продолжим поиск
		} else if val < s {
			left = mid + 1
			// Аналогично в противоположной ситуации
		} else {
			right = mid - 1

		}
	}
	// Если необходимого значения в массиве нет вернем -1
	return -1
}

func main() {
	arr := quicksort(generateSlice(20))
	for idx, v := range arr {
		fmt.Println(idx, v)
	}

	val := arr[rand.Int()%len(arr)]
	fmt.Println("Ищем индекс значения", val)
	fmt.Println(binarySerach(arr, val))
	impossibleVal := 333 // Максимальное возможное значение массива - 299
	fmt.Println("Ищем индекс значения", impossibleVal)
	fmt.Println(binarySerach(arr, impossibleVal))
}
