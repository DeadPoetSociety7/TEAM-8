# team8
Interactive Kanban CLI

A lightweight, interactive command-line Kanban board built entirely with Go's standard library. This application allows teams and individuals to manage tasks directly from the terminal, keeping data persistent across runs. 

Progress Tracking: Tasks move through three sequential statuses: todo ➔ inprogress ➔ done. 


Data Persistence: Automatically saves and loads tasks to/from a flat file database. 


Input Validation: Built-in protection against invalid IDs, empty strings, and bad menu choices to prevent application crashes. 


Russian Language Support: Beautiful terminal table auto-aligns properly even when task titles contain Cyrillic characters. 


Requirements
Go 1.16 or higher 


"bufio","fmt","os","strconv","strings"libraries used

How to Run
Interactive Menu
To start the interactive Kanban board application, navigate to the project directory and run: 

bash
go run main.go
Use code with caution.

Help Command
To view usage information and available arguments: 

bash
go run main.go --help
Use code with caution.

System Rules & Data Format
Menu Actions
Create Task: Prompts for a title, generates a unique ID (max_id + 1), and initializes the status as todo. 


Advance Status: Moves a task forward step-by-step (todo ➔ inprogress ➔ done). 


Delete Task: Removes a task permanently from the system by its ID.

All Tasks: Displays a beautifully formatted board with all tasks. 


Edit Title: Renames an existing task title by its ID. 

Exit: Safely terminates the program execution. 

Storage Format
All tasks are persisted in the root directory inside tasks.txt. The data is structured using a semicolon-separated line format: 

text
id;status;title
Use code with caution.

Example contents of tasks.txt: 

text
1;todo;Setup project
2;inprogress;Write code
3;done;Review pull request
Use code with caution.

Note: Any malformed lines containing fewer than 3 fields are skipped on load. 

Example Session
text
== МЕНЮ ЗАДАЧ ==
1. Создать таску
2. Изменить статус задачи
3. Удалить таску
4. Все таски
5. Завершить работу программы
Выберите действие: 4

+----+----------------------+-------------+

| ID | Title                | Status      |
+----+----------------------+-------------+

|  1 | Setup project        | todo        |
|  2 | Написать код         | inprogress  |
|  3 | Review               | done        |
+----+----------------------+-------------+
Use code with caution.

Team Members:
Member 1 - Grond3749

Member 2 - Musenova-lgtm

Member 3 - DeadPoetSociety7
