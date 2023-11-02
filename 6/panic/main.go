package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		// Эта горутина вызывает панику, выполняя деление на ноль
		a := 1
		b := 0
		result := a / b     // Это вызовет панику
		fmt.Println(result) // Этот код не будет выполнен из-за паники
	}()

	fmt.Println("Главная горутина продолжает работу")
	// Ждем некоторое время, чтобы горутина имела возможность выполниться
	wg.Wait()
}
