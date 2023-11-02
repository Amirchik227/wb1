package main

import (
	"fmt"
	"sync"
)

// Находим квадрат числа
func squaringNumber(wg *sync.WaitGroup, ch chan int, number int) {
	// Перед завершением squaringNumber сообщаем sync.WaitGroup,
	// что очередная задача выполнена
	defer wg.Done()
	// Кладем в канал квадрат числа
	ch <- number * number

}

func main() {
	// sync.WaitGroup необходим для ожидания завершения всех горутин
	var wg sync.WaitGroup

	numbers := []int{2, 4, 6, 8, 10}

	// Канал, по которому передаются квадраты чисел
	squaredNumbersChan := make(chan int)
	defer close(squaredNumbersChan) // После выполнения программы канал закроется

	// Для каждого числа создаем отдельную горутину, которая находит его квадарат
	for _, number := range numbers {
		// Добавляем 1 процесс в sync.WaitGroup, чтобы main не завершился,
		// пока мы его не посчитали очередной квадрат
		wg.Add(1)
		go squaringNumber(&wg, squaredNumbersChan, number)
	}

	wg.Add(1)
	// В отдельной горутине считываем квадраты из канала и суммируем
	go func() {
		defer wg.Done()
		sum := 0
		for i := 0; i < len(numbers); i++ {
			sum += <-squaredNumbersChan
		}
		fmt.Println(sum)
	}()

	wg.Wait()
}
