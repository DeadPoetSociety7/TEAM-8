package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename = "tasks.txt"

type Task struct {
	ID     int
	Title  string
	Status string
}

var tasks []Task

func main() {
	if len(os.Args) > 1 {
		arg := os.Args[1]
		if arg == "-h" || arg == "--help" {
			print_help()
			return
		}
		fmt.Println("Неизвестный аргумент. Используйте --help для справки.")
		os.Exit(1)
	}

	load_tasks()
	menu()
}

func print_help() {
	fmt.Println("Kanban CLI — консольный менеджер задач в интерактивном режиме.")
	fmt.Println()
	fmt.Println("Использование:")
	fmt.Println("  go run main.go         Запуск интерактивного меню")
	fmt.Println("  go run main.go --help  Показать эту справку")
	fmt.Println()
	fmt.Println("Доступные команды в меню:")
	fmt.Println("  1. Создать таску          — запрашивает название и добавляет новую задачу (TODO)")
	fmt.Println("  2. Пометить выполненным   — переводит задачу по ID в статус DONE")
	fmt.Println("  3. Удалить таску          — удаляет задачу из списка по ID")
	fmt.Println("  4. Все таски              — выводит текущую доску задач в виде ASCII-таблицы")
	fmt.Println("  5. Завершить работу       — сохраняет состояние и закрывает программу")
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		print_menu()
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

			if title == "" {
				fmt.Println("Error: название не может быть пустым.")
				continue
			}

			id := max_ID() + 1
			tasks = append(tasks, Task{ID: id, Title: title, Status: "TODO"})
			save_tasks()

		case "2":
			fmt.Print("Введите ID задачи для отметки: ")
			if !scanner.Scan() {
				continue
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: ID должен быть числом.")
				continue
			}

			idx := findindexbyID(id)
			if idx == -1 {
				fmt.Printf("Error: задача с ID %d не найдена.\n", id)
				continue
			}

			if tasks[idx].Status == "DONE" {
				fmt.Printf("Error: задача %d уже выполнена.\n", id)
				continue
			}

			tasks[idx].Status = "DONE"
			save_tasks()

		case "3":
			fmt.Print("Введите ID задачи для удаления: ")
			if !scanner.Scan() {
				continue
			}
			id, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Error: ID должен быть числом.")
				continue
			}

			idx := findindexbyID(id)
			if idx == -1 {
				fmt.Printf("Error: задача с ID %d не найдена.\n", id)
				continue
			}

			tasks = append(tasks[:idx], tasks[idx+1:]...)
			save_tasks()

		case "4":
			print_tasks(tasks)

		case "5":
			os.Exit(0)

		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 5.")
		}
	}
}

func print_menu() {
	fmt.Println("\n=== МЕНЮ ЗАДАЧ ===")
	fmt.Println("1. Создать таску")
	fmt.Println("2. Пометить выполненным")
	fmt.Println("3. Удалить таску")
	fmt.Println("4. Все таски")
	fmt.Println("5. Завершить работу программы")
	fmt.Print("Выберите действие (1-5): ")
}

func print_tasks(tasks_list []Task) {
	if len(tasks_list) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Println("\n+----+----------------------+-------------+")
	fmt.Println("| ID | Title                | Status      |")
	fmt.Println("+----+----------------------+-------------+")

	for _, task := range tasks_list {
		runes := []rune(task.Title)
		display_title := task.Title

		if len(runes) > 20 {
			display_title = string(runes[:17]) + "..."
			runes = []rune(display_title)
		}

		spaces_cnt := 20 - len(runes)
		spaces := ""
		if spaces_cnt > 0 {
			spaces = strings.Repeat(" ", spaces_cnt)
		}

		fmt.Printf("| %2d | %s%s | %-11s |\n", task.ID, display_title, spaces, task.Status)
	}
	fmt.Println("+----+----------------------+-------------+")
}

func max_ID() int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max
}

func findindexbyID(id int) int {
	for i, t := range tasks {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func load_tasks() {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		parts := strings.Split(line, ";")
		if len(parts) != 3 {
			continue
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			continue
		}

		status := strings.ToUpper(strings.TrimSpace(parts[1]))
		if status != "TODO" && status != "DONE" {
			continue
		}

		title := strings.TrimSpace(parts[2])
		tasks = append(tasks, Task{ID: id, Title: title, Status: status})
	}
}

func save_tasks() {
	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, t := range tasks {
		writer.WriteString(fmt.Sprintf("%d;%s;%s\n", t.ID, t.Status, t.Title))
	}
	writer.Flush()
}
