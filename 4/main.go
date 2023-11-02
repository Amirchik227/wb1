package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Вводим количество ворекров
	var workersCount int
	fmt.Println("Enter the number of workers")
	fmt.Scan(&workersCount)
	// Создаем контекст и функцию для его завершения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)     // Со-даем канал, котрый будет заврешать работу
	signal.Notify(sigChan, syscall.SIGINT) // программы по нажантию Ctrl+C.

	go func() { // В отдельной горутине
		<-sigChan // ожидаем сигнала о завершении работы программы,
		cancel()  // а после закрываем канал.
	}()

	randomChan := make(chan int) // Канал по передаче произвольных данных
	var wg sync.WaitGroup        // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1) // Увеличение счетчика процессов, завершения которых мы будем ожидать
	go func() {
		defer wg.Done()         // Уменьшение счётчика по завершении выполнения
		writer(ctx, randomChan) //  Начинаем запись данных в канал
	}()

	for i := 0; i < workersCount; i++ {
		wg.Add(1)        // Для каждого воркера инкрементируем счётчик
		go func(i int) { // В отдельной горутине запускаем i-го воркера
			defer wg.Done()            // Не забваем уменьшить счетчик после завершения работы воркера
			worker(ctx, randomChan, i) // Запуск i-го воркера
		}(i) // Передаем номер воркера
	}
	wg.Wait() // Ожидаем завершения всех процессов
}

// writer заполняет канал случайными данными
func writer(ctx context.Context, ch chan int) {
	defer close(ch) // Не заабываем закрыть какнал
	// Инициализация генератора случайных чисел
	var rg = rand.New(rand.NewSource(time.Now().Unix()))
	//Создадим ticke, чтобы данные не генерировались слишком быстро
	ticker := time.NewTicker(150 * time.Millisecond)
	for {
		select {
		// Если получен сигнал о завершении программы перестаем заполнять канал
		case <-ctx.Done():
			fmt.Println("Program stopped")
			return
		case <-ticker.C: // Если верям пришло,
			data := rg.Intn(9000) + 1000 // сгенерируем число от 1000 до 9999
			ch <- data                   // и передадим в канал
		}
	}
}

// Воркер считывает данные из канала и выводит в stdout
func worker(ctx context.Context, ch chan int, i int) {
	for {
		select {
		// Если получен сигнал о завершении программы, завершаем работу воркера
		case <-ctx.Done():
			return
		default:
			// Пробуем считать значение из канала
			value, ok := <-ch
			// Если чтение прошло успешно, выводим полученные данные
			if ok {
				fmt.Printf("Worker №%d: %d\n", i, value)
			}
		}
	}
}
