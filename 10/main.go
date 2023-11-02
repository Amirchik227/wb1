package main

import (
	"fmt"
	"strconv"
)

func main() {
	temperatures := [...]float64{-25.4, -29.99, -20.01 - 27.0, 13.0, 19.0, 15.5,
		24.5, -21.0, 32.5, 5.0, 0.5, 0.6, 0.0, -0.7, -5, -0.9, -10.5, -10}

	// Карта,  в которой будут группы температур
	groups := make(map[string][]float64)

	for _, temp := range temperatures {
		// Так как шаг группы 10, отставим в числе только десятки и более
		// старшие разряды, обнулив единицы
		group := strconv.Itoa(int(temp) / 10 * 10)
		// Температуры в промежутках (-10;0) и [0;10) необходимо разбить в разные группы
		// Выделим промежуток (-10;0) от отдельную группу "-0"
		if temp > -10 && temp < 0 {
			group = "-0"
		}
		//Добавим температуру в свою группу
		groups[group] = append(groups[group], temp)
	}
	// Вывод групп
	for group, values := range groups {
		fmt.Printf("%s : %v\n", group, values)
	}
}
