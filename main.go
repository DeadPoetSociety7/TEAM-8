package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID     int
	Title  string
	Status string
}

var tasks []Task

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "-h" || arg == "--h" || arg == "--help" {
			printHelp()
			return
		} else {
			fmt.Println("Неизвестный аргумент. Используйте --help для справки.")
			os.Exit(1)
		}
	}

	// Начальные данные для демонстрации
	tasks = []Task{
		{1, "Setup project", "todo"},
		{2, "Write code", "inprogress"},
		{3, "Review", "done"},
	}
	menu()
}

func printHelp() {
	fmt.Println("Interactive Kanban CLI")
	fmt.Println()
	fmt.Println("Использование: go run main.go [флаг]")
	fmt.Println()
	fmt.Println("Флаги:")
	fmt.Println("  -h, --h, --help   Показать эту справку и выйти")
	fmt.Println()
	fmt.Println("Без флагов запускается интерактивное меню:")
	fmt.Println("  1. Создать таску")
	fmt.Println("  2. Пометить выполненным")
	fmt.Println("  3. Удалить таску")
	fmt.Println("  4. Все таски")
	fmt.Println("  5. Завершить работу программы")
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	runMenu(scanner)
}

func runMenu(scanner *bufio.Scanner) {
	for {
		printMenu()
		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Введите название задачи: ")
			if !scanner.Scan() {
				continue
			}
			title := strings.TrimSpace(scanner.Text())

			// Валидация: название не должно быть пустым
			if title == "" {
				fmt.Println("Ошибка: Название задачи не может быть пустым!")
				continue
			}

			id := len(tasks) + 1
			tasks = append(tasks, Task{id, title, "todo"})
			fmt.Println("Задача успешно создана.")

		case "2":
			fmt.Print("Введите ID задачи для отметки: ")
			if !scanner.Scan() {
				continue
			}
			input := strings.TrimSpace(scanner.Text())

			// Валидация: перевод строки в число
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом!")
				continue
			}

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Status = "done"
					fmt.Println("Задача отмечена как выполненная.")
					found = true
					break
				}
			}

			// Валидация: проверка, найден ли ID в списке
			if !found {
				fmt.Printf("Ошибка: Задача с ID %d не найдена.\n", id)
			}

		case "3":
			fmt.Print("Введите ID задачи для удаления: ")
			if !scanner.Scan() {
				continue
			}
			input := strings.TrimSpace(scanner.Text())

			// Валидация: перевод строки в число
			id, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Ошибка: ID должен быть числом!")
				continue
			}

			found := false
			for i := range tasks {
				if tasks[i].ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Задача успешно удалена.")
					found = true
					break
				}
			}

			// Валидация: проверка, найден ли ID в списке
			if !found {
				fmt.Printf("Ошибка: Задача с ID %d не найдена.\n", id)
			}

		case "4":
			printTasks(tasks)

		case "5":
			fmt.Println("Программа завершена.")
			os.Exit(0)

		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 5.")
		}
	}
}

func printMenu() {
	fmt.Println("\n== МЕНЮ ЗАДАЧ ==")
	fmt.Println("1. Создать таску")
	fmt.Println("2. Пометить выполненным")
	fmt.Println("3. Удалить таску")
	fmt.Println("4. Все таски")
	fmt.Println("5. Завершить работу программы")
	fmt.Print("Выберите действие: ")
}

func printTasks(tasksList []Task) {
	if len(tasksList) == 0 {
		fmt.Println("\nСписок задач пуст.")
		return
	}

	fmt.Println("\n+----+----------------------+-------------+")
	fmt.Println("| ID | Title                | Status      |")
	fmt.Println("+----+----------------------+-------------+")

	for _, task := range tasksList {
		// Преобразуем в срез рун для корректной работы с русскими буквами
		runes := []rune(task.Title)
		displayTitle := task.Title

		// Ограничиваем длину заголовка до 20 видимых символов
		if len(runes) > 20 {
			displayTitle = string(runes[:17]) + "..."
			runes = []rune(displayTitle)
		}

		// Вычисляем, сколько пробелов нужно добавить для выравнивания (всего 20 мест)
		spacesCount := 20 - len(runes)
		spaces := ""
		if spacesCount > 0 {
			spaces = strings.Repeat(" ", spacesCount)
		}

		// Форматируем статус (он на английском, тут стандартные 11 мест)
		fmt.Printf("| %2d | %s%s | %-11s |\n", task.ID, displayTitle, spaces, task.Status)
	}
	fmt.Println("+----+----------------------+-------------+")
}