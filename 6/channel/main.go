package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	quitChan := make(chan int) // Канал по передаче произвольных данных
	defer close(quitChan)
	var wg sync.WaitGroup // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1)          // Для воркера инкрементируем счётчик
	go worker(quitChan, &wg) // Запуск воркера
	time.Sleep(1 * time.Second)
	quitChan <- 1
	wg.Wait() // Ожидаем завершения всех процессов
}

func worker(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	//Создадим ticker, чтобы данные не генерировались слишком быстро
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		// Если получен сигнал о завершении программы перестаем заполнять канал
		case <-ch:
			fmt.Printf("Горутина остановлена через канал")
			return
		case <-ticker.C: // Если верям пришло,
			fmt.Println(time.Now())
		}
	}
}
