package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup // Создаем sync.WaitGroup для ожидания завершения всех горутин
	isWorking := true
	wg.Add(1) // Для воркера инкрементируем счётчик
	go worker(&wg, &isWorking)
	time.Sleep(700 * time.Millisecond)
	isWorking = false
	wg.Wait() // Ожидаем завершения всех процессов
}

func worker(wg *sync.WaitGroup, isWorking *bool) {
	defer wg.Done()
	//Создадим ticker, чтобы данные не генерировались слишком быстро
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		if !(*isWorking) {
			fmt.Println("Работа горутины остановлена, так как изменилось значение на которое указывает isWorking")
			return
		}
		_, ok := <-ticker.C
		if ok {
			fmt.Println(time.Now())
		}
	}
}
