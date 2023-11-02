package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	// Ожидаем получения значения из канала, возвращаемое из time.After
	// Записывать значение нет необходимости
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println(time.Now())
	// Приостанавливаем выполнение на 2 секунды
	sleep(2)
	fmt.Println(time.Now())
}
