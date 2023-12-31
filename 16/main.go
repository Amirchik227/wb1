package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создаем массив из 20 случайных чисел
	slice := generateSlice(20)
	fmt.Println("Неотсортированный массив\n", slice)
	quicksort(slice)
	fmt.Println("Отсортированный массив\n", slice)
}

func generateSlice(size int) []int {
	// Создаем слайс из size дефолтных значений
	slice := make([]int, size)
	// Переменная для генерации случайных чисел
	rg := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < size; i++ {
		// Заполняем слайс случайными значениями от 0 до 299
		slice[i] = rg.Intn(300)
	}
	return slice
}

func quicksort(a []int) []int {
	// Quicksort решает задачу сортировки методом разделяй и властвуй
	// Если после разделения в нашу подзадачу попал массив размер которого
	// меньше двух, то он не нуждается в сортировке
	if len(a) < 2 {
		return a
	}
	// Если же его длина > 1, то начнем с определения границ массива
	left, right := 0, len(a)-1

	// Случайно определим элемент массива pivot, а в дальнейшем проведем
	// сортировку таким образом, чтобы всле элементы массива слева от него
	// обладали меньшим значением, а все элементы справа - бОльшим или равным
	pivot := rand.Int() % len(a)

	// Для удобства сортировки поменяем значение крайнего правого элемента и
	// pivot элемента. Во премя сортировки все элементы больше pivot будут
	// приближаться к правому краю массива, а все элменты меньше - к левому
	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		// Если a[i] меньше a[right](по сути pivot), то можной спокойно
		// отправить его в левый край и забыть о нем, сдвинув левую границу.
		// Таким образом постоянно смещая a[left] можно быть уверенным, что
		// все значения слева от a[left] будут меньше pivot.
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
		// В конце работы алгоритма в a[left] будет знчаение большее или равное
		// pivot. Так как pivot является наименьшим знаением из тех, чей
		// индекс >= left(так как все значения меньше его мы переместили влево),
		// то можно пославить его в начало подмассива a[left:right+1]
	}

	a[left], a[right] = a[right], a[left]
	// Далее продолжаем разделять и властвовать.
	// pivot уже стоит на своем месте (все значения слева от него меньше его,
	// а справа - больше), поэтому применяем тот же алгоритм к двум подмассивам
	// слева и справа от pivot. Рекурсивно продолжная данный алгоритм отсортируем
	// весь массив
	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
