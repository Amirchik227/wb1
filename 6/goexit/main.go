package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1) // Для воркера инкрементируем счётчик
	go worker(&wg, 1)
	wg.Add(1) // Для воркера инкрементируем счётчик
	go worker(&wg, 0)

	wg.Wait() // Ожидаем завершения всех процессов
}

func worker(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	if i == 0 {
		fmt.Println("Горутина остановлена Goexit")
		runtime.Goexit()
	}
	bigTicker := time.NewTicker(1 * time.Second)
	smallTicker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-bigTicker.C:
			fmt.Printf("Горутина остановлена по истечении одной секунды")
			return
		case <-smallTicker.C:
			fmt.Printf("%d: %s\n", i, time.Now())
		}
	}
}
