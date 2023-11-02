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
	bigTicker := time.NewTicker(1 * time.Second)
	smallTicker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-bigTicker.C:
			fmt.Printf("Горутина остановлена по истечении одной секунды")
			return
		case <-smallTicker.C:
			fmt.Println(time.Now())
		}
	}
}
