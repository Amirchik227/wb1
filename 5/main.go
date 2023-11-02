package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// Вводим количество секунд
	var duration int
	fmt.Println("Enter the number of seconds")
	fmt.Scan(&duration)
	// Создаем контекст, который  и функцию для его завершения
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()
	randomChan := make(chan int) // Канал по передаче произвольных данных
	var wg sync.WaitGroup        // Создаем sync.WaitGroup для ожидания завершения всех горутин

	wg.Add(1) // Увеличение счетчика процессов, завершения которых мы будем ожидать
	go func() {
		defer wg.Done()         // Уменьшение счётчика по завершении выполнения
		writer(ctx, randomChan) //  Начинаем запись данных в канал
	}()

	wg.Add(1)   // Для воркера инкрементируем счётчик
	go func() { // В отдельной горутине запускаем воркера
		defer wg.Done()         // Не забываем уменьшить счетчик после завершения работы воркера
		worker(ctx, randomChan) // Запуск воркера
	}()
	wg.Wait() // Ожидаем завершения всех процессов
}

// writer заполняет канал случайными данными
func writer(ctx context.Context, ch chan int) {
	defer close(ch) // Не заабываем закрыть какнал
	// Инициализация генератора случайных чисел
	var rg = rand.New(rand.NewSource(time.Now().Unix()))
	//Создадим ticker, чтобы данные не генерировались слишком быстро
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		// Если получен сигнал о завершении программы перестаем заполнять канал
		case <-ctx.Done():
			return
		case <-ticker.C: // Если верям пришло,
			data := rg.Intn(9000) + 1000 // сгенерируем число от 1000 до 9999
			ch <- data                   // и передадим в канал
		}
	}
}

// Воркер считывает данные из канала и выводит в stdout
func worker(ctx context.Context, ch chan int) {
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
				fmt.Printf("%d\n", value)
			}
		}
	}
}
