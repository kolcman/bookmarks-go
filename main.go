package main

import "fmt"

var bookmarks = map[string]string{}

func main() {
	for {
		showMenu()
		handleChoice()
	}
}

func showMenu() {
	fmt.Print(`
*--------------------------*
|       МЕНЮ ЗАКЛАДОК      |
*--------------------------*
| 1. Посмотреть закладки   |
| 2. Добавить закладку     |
| 3. Удалить закладку      |
| q. Выход                 |
*--------------------------*
Ваш выбор: `)
}

func handleChoice() {
	var choice string
	fmt.Scan(&choice)
	switch choice {
	case "1":
		showBookmarks()
	case "2":
		addBokmark()
	case "3":
	// removeBokmark()
	case "q", "Q", "й", "Й":
		return
	}
}

func showBookmarks() {
	for k, v := range bookmarks {
		fmt.Printf("Название: %s - адресс: %s\n\n\n", k, v)
	}
}

func addBokmark() {

}
