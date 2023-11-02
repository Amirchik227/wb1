package main

import (
	"fmt"
)

// Имеем структуру, которая может сгенерировать пароль
type AdapteeGen struct{}

// Метод генерации пароля
func (a *AdapteeGen) GeneratePass() string {
	return "qwerty123"
}

// Необходима структура, реализующая интерфейс Generator
type Generator interface {
	GeneratePassword() string
}

// У нас уже есть структура (AdapteeGen), сопсобная генерировать пароль
// Адаптируем её под данный интерфейс
type Adapter struct {
	*AdapteeGen
}

// Адаптируем метод GeneratePass структуры AdapteeGen
// Чтобы он подходил под наш интерфейс
func (adapter *Adapter) GeneratePassword() string {
	str := adapter.GeneratePass()
	return str
}

// Конструктор адаптера
func NewAdapter(adaptee *AdapteeGen) Generator {
	return &Adapter{adaptee}
}

func main() {
	// Инициализия адаптера
	adaptedGen := NewAdapter(&AdapteeGen{})
	// Использование адаптера
	fmt.Println(adaptedGen.GeneratePassword())
}
