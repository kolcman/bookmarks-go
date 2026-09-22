package main

import "fmt"

var bookmarks = map[string]string{}

func main() {
	for {
		showMenu()
		var choice string
		fmt.Scan(&choice)
		if choice == "q" {
			break
		}
		handleChoice(choice)
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

func handleChoice(choice string) {
	switch choice {
	case "1":
		showBookmarks()
	case "2":
		addBokmark()
	case "3":
	// removeBokmark()
	default:
		fmt.Println("Неккоректный ввод!")
	}
}

func showBookmarks() {
	for k, v := range bookmarks {
		fmt.Printf("Название: %s - адресс: %s\n", k, v)
	}
	fmt.Print("\n\n\n")
}

func addBokmark() {
	name := getValue("имя")
	url := getValue("адрес")
	bookmarks[name] = url
}

func getValue(typeValue string) string {
	var input string
	fmt.Printf("Введите %s закладки: ", typeValue)
	fmt.Scan(&input)
	return input
}
