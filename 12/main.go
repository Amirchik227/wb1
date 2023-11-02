package main

import "fmt"

func main() {
	var array = []string{"cat", "cat", "dog", "cat", "tree"}
	var set = []string{}
	// Карта будет хранить в себе в качестве ключей все значения массива
	var presenceMap = make(map[string]bool)

	// Заполнение карты
	for _, item := range array {
        presenceMap[item] = true
    }
	// Заполнение сета уникальными ключами карты
	for key := range presenceMap {
		set = append(set, key)
		fmt.Println(key)
	}
}
