package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// Создаем контекст, который закроется автоматически через 1 сек
	// и функцию для его завершения
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1) // Для воркера инкрементируем счётчик
	go worker(ctx, &wg)

	wg.Wait() // Ожидаем завершения всех процессов
}

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	//Создадим ticker, чтобы данные не генерировались слишком быстро
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		// Если получен сигнал о завершении программы перестаем заполнять канал
		case <-ctx.Done():
			fmt.Printf("Горутина остановлена ограниченным по времени контекстом")
			return
		case <-ticker.C: // Если верям пришло,
			fmt.Println(time.Now())
		}
	}
}
