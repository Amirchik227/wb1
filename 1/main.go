package main

import "fmt"

// Определим структуру Human
type Human struct {
	Name   string
	Height uint
}

// Метод для структуры Human
func (h *Human) Introduction() {
	fmt.Printf("Привет, меня зовут %s, мой рост %d см.\n", h.Name, h.Height)
}

// Определим структуру Action и встроим в неё структуру Human
type Action struct {
	Human
}

// Метод для структуры Action
// func (a *Action) DoAction() {
//     fmt.Printf("%s выполняет действие: %s\n", a.Name, a.ActionName)
// }

func main() {
	// Создадим экземпляр структуры Action
	action := Action{
		Human: Human{Name: "Майкл", Height: 198},
	}

	// Вызовем метод Introduction,
	// который будет ближайшим к корню структуры Action
	action.Introduction()
	// Обратимся к напрамую к методу Introduction структуры Human
	action.Human.Introduction()
}
