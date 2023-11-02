package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Не понятно зачем использовать глобальную переменную,
// Выглядит так, как будто можно ограничиться локальной
//var justString string

// Так как в коде подразумевается использование только 100 символьных строк
// можно ограничить size или или выдать ошибку
// if size > 100 {
// 	size = 100
// }

func createHugeString(size int) (string, error) {
	// Если размер слишком большой возвращаем ошибку
	if size > 100 {
		return "", errors.New("СЛИШКОМ БОЛЬШОЙ РАЗМЕР СТРОКИ")
	}
	// Переменная для генерации случайного числа
	rg := rand.New(rand.NewSource(time.Now().Unix()))
	// Символьная переменная
	var c rune
	// Создадим слайс, у которого длина 0 и емкость size
	result := make([]rune, 0, size)
	for i := 0; i < size; i++ {
		// Сгенерируем ascii код японских символов
		c = rune(12352 + rg.Intn(192)) // Японский алфавит
		result = append(result, c)
	}
	return string(result), nil
}

func someFunc() {
	//v, err := createHugeString(1 << 10)
	justString, err := createHugeString(50)
	if err != nil {
		fmt.Println("error creating string:", err)
	}
	// В переменную justString передаются первые 100 байтов из v,
	// а не первые 100 символов как можно было бы ожидать. В случае с японскими
	// символами это может привести к ошибке, ведь они весят больше 1 байта
	//justString = v[:100]

	// Чтобы это исправить это переведем строку в слайс рун (тут происходит индексация
	// по символам, а не по байтам) и обратно
	justString = string([]rune(justString)[:25])
	fmt.Printf("%s\n", justString)
}

func main() {
	someFunc()
}
