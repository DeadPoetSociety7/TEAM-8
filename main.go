package main

import (
	
)

type Task struct {
	ID     int
	Title  string
	Status string
}

var tasks []Task

const filename = "tasks.txt"

func main() {
	// Проверяем аргументы командной строки
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			if arg == "-h" || arg == "--help" {
				printHelp()
				return
			}

			fmt.Println("Неизвестный аргумент. Используйте --help для справки.")
			return
		}
	}

	// Загружаем задачи из файла
	loadTasks()

	// Запускаем меню
	menu()
}

func printHelp() {
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
	fmt.Println("  4. Все таски              — выводит текущую доску задач")
	fmt.Println("  5. Завершить работу       — сохраняет состояние и закрывает программу")
}

func menu() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMenu()

		if !scanner.Scan() {
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			createTask(scanner)

		case "2":
			markDone(scanner)

		case "3":
			deleteTask(scanner)

		case "4":
			printTasks(tasks)

		case "5":
			saveTasks()
			fmt.Println("Программа завершена.")
			return

		default:
			fmt.Println("Неверный ввод. Введите число от 1 до 5.")
		}
	}
}

func printMenu() {
	fmt.Println()
	fmt.Println("=== МЕНЮ ЗАДАЧ ===")
	fmt.Println("1. Создать таску")
	fmt.Println("2. Пометить выполненным")
	fmt.Println("3. Удалить таску")
	fmt.Println("4. Все таски")
	fmt.Println("5. Завершить работу программы")
	fmt.Print("Выберите действие (1-5): ")
}

func createTask(scanner *bufio.Scanner) {
	fmt.Print("Введите название задачи: ")

	if !scanner.Scan() {
		return
	}

	title := strings.TrimSpace(scanner.Text())

	if title == "" {
		fmt.Println("Error: название не может быть пустым.")
		return
	}

	id := maxID() + 1

	tasks = append(tasks, Task{
		ID:     id,
		Title:  title,
		Status: "TODO",
	})

	saveTasks()

	fmt.Println("Задача успешно создана.")
}

func markDone(scanner *bufio.Scanner) {
	fmt.Print("Введите ID задачи: ")

	if !scanner.Scan() {
		return
	}

	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: ID должен быть числом.")
		return
	}

	index := findIndexByID(id)

	if index == -1 {
		fmt.Printf("Error: задача с ID %d не найдена.\n", id)
		return
	}

	if tasks[index].Status == "DONE" {
		fmt.Printf("Error: задача %d уже выполнена.\n", id)
		return
	}

	tasks[index].Status = "DONE"

	saveTasks()

	fmt.Println("Задача отмечена как выполненная.")
}

func deleteTask(scanner *bufio.Scanner) {
	fmt.Print("Введите ID задачи: ")

	if !scanner.Scan() {
		return
	}

	input := strings.TrimSpace(scanner.Text())

	id, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Error: ID должен быть числом.")
		return
	}

	index := findIndexByID(id)

	if index == -1 {
		fmt.Printf("Error: задача с ID %d не найдена.\n", id)
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	saveTasks()

	fmt.Println("Задача успешно удалена.")
}

func printTasks(tasksList []Task) {
	if len(tasksList) == 0 {
		fmt.Println("Список задач пуст.")
		return
	}

	fmt.Println()
	fmt.Println("+----+----------------------+-------------+")
	fmt.Println("| ID | Title                | Status      |")
	fmt.Println("+----+----------------------+-------------+")

	for _, task := range tasksList {
		title := task.Title

		runes := []rune(title)

		if len(runes) > 20 {
			title = string(runes[:17]) + "..."
			runes = []rune(title)
		}

		spaces := 20 - len(runes)

		if spaces < 0 {
			spaces = 0
		}

		fmt.Printf(
			"| %2d | %s%s | %-11s |\n",
			task.ID,
			title,
			strings.Repeat(" ", spaces),
			task.Status,
		)
	}

	fmt.Println("+----+----------------------+-------------+")
}

func maxID() int {
	max := 0

	for _, task := range tasks {
		if task.ID > max {
			max = task.ID
		}
	}

	return max
}

func findIndexByID(id int) int {
	for i, task := range tasks {
		if task.ID == id {
			return i
		}
	}

	return -1
}

func loadTasks() {
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

		parts := strings.SplitN(line, ";", 3)

		if len(parts) != 3 {
			continue
		}

		id, err := strconv.Atoi(strings.TrimSpace(parts[0]))

		if err != nil {
			continue
		}

		status := strings.ToUpper(strings.TrimSpace(parts[1]))

		if status != "TODO" && status != "DONE" {
			continue
		}

		title := strings.TrimSpace(parts[2])

		tasks = append(tasks, Task{
			ID:     id,
			Title:  title,
			Status: status,
		})
	}
}

func saveTasks() {
	file, err := os.Create(filename)

	if err != nil {
		fmt.Println("Error: не удалось сохранить задачи.")
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, task := range tasks {
		writer.WriteString(
			fmt.Sprintf(
				"%d;%s;%s\n",
				task.ID,
				task.Status,
				task.Title,
			),
		)
	}

	writer.Flush()
}     