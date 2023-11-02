package main

import (
	"fmt"
	"reflect"
)

func typeOf(x interface{}) {
	// Используем switch для перебора возмождных вариантов значения переменой
	switch x.(type) {
	case int:
		fmt.Printf("%v : int\n", x)
	case string:
		fmt.Printf("%v : string\n", x)
	case bool:
		fmt.Printf("%v : bool\n", x)
	case chan int:
		fmt.Printf("%v : chan int\n", x)
	case chan string:
		fmt.Printf("%v : chan string\n", x)
	default:
		fmt.Printf("%v : неизвестен\n", x)
	}
}

func main() {
	array := []interface{}{"text", 2344, true, make(chan int)}
	for _, v := range array {
		typeOf(v)
		// Используем библиотеку reflect для определения типа
		fmt.Printf("%v : %v\n", v, reflect.TypeOf(v))
		// Также можно использовать библиотеку "fmt"
		fmt.Printf("%v : %T\n", v, v)
	}
}
