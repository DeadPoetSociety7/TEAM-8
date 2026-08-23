package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("=== МЕНЮ ЗАДАЧ ===")
		fmt.Println("1. Создать таску")
		fmt.Println("2. Пометить выполненным")
		fmt.Println("3. Удалить таску")
		fmt.Println("4. Все таски")
		fmt.Println("5. Завершить работу программы")
		fmt.Println("Выберите действие (1-5):")

		if !scanner.Scan() {
			break
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Println("Создание задач")
		case "2":
			fmt.Println("Выполнение задачи")
		case "3":
			fmt.Println("Удаление задачи")
		case "4":
			fmt.Println("Вывод всех задач")
		case "5":
			return
		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 5.")
		}
	}
}

func main() {
	menu()
}
