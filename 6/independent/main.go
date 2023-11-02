package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1) // Для воркера инкрементируем счётчик
	go worker(&wg)

	wg.Wait() // Ожидаем завершения всех процессов
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 7; i++ {
		fmt.Println(time.Now())
		time.Sleep(1 * time.Millisecond)
	}
	fmt.Printf("Горутина остановлена по завершению выполнения тела функции")
}
