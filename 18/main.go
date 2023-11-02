package main

import (
	"fmt"
	"sync"
)

// Структура, в которой хранится счетчик
type CounterStorage struct {
	// sync.RWMutex  для блокировки и разблокировки доступа к счетчику в конкурентной среде
	mx    sync.RWMutex 
	// Сам счетчик
	counter int          
}

// Функция инициализации нового CounterStorage
func NewCounterStorage() *CounterStorage {
	return &CounterStorage{
		counter: 0,
	}
}

// Инкремент счетчика
func (c *CounterStorage) Inc() {
	// Блокируем запись для того, чтобы две горутины не увеличивали 
	// счетчик одновременно. Иначе может возникнуть ситуация, когда две
	// горутины получат одно и то же значение counter и совместно увеличат
	// его не на 2, а на 1.
	c.mx.Lock()         
	defer c.mx.Unlock() // Разблокируем запись после инкремента
	c.counter++ // Инкремент
}

// Получение значения счетчика
func (c *CounterStorage) GetCounter() int {
	c.mx.RLock()         // Блокируем чтение
	defer c.mx.RUnlock() // Разблокируем чтение
	return c.counter // Возвращаем значение
}

func main() {
	// Инициализация счетчика
	var counter = NewCounterStorage()

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for i := 0; i < 1000001; i++ {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go func() {
			defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
			counter.Inc()
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результата
	fmt.Println(counter.GetCounter())

}
