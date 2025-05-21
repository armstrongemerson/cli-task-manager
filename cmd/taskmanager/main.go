package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "github.com/yourusername/taskmanager/internal/task"
)

func main() {
	// Initialize task repository
	repo, err := task.NewRepository("tasks.json")
	if err != nil {
		fmt.Println("Error initializing task repository:", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("============")
		fmt.Println("1. List tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. View task details")
		fmt.Println("5. Save and exit")
		fmt.Print("\nChoose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			listTasks(repo)
		case "2":
			addTask(repo, scanner)
		case "3":
			markTaskAsDone(repo, scanner)
		case "4":
			viewTaskDetails(repo, scanner)
		case "5":
			if err := repo.Save(); err != nil {
				fmt.Println("Error saving tasks:", err)
			} else {
				fmt.Println("Tasks saved successfully.")
			}
			return
		default:
			fmt.Println("Invalid option, please try again.")
		}
	}
}

func listTasks(repo *task.Repository) {
	tasks := repo.List()

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("\nYour Tasks:")
	for _, t := range tasks {
		fmt.Println(t.Summary())
	}
}

func addTask(repo *task.Repository, scanner *bufio.Scanner) {
	fmt.Print("Enter task title: ")
	scanner.Scan()
	title := scanner.Text()

	fmt.Print("Enter task description: ")
	scanner.Scan()
	description := scanner.Text()

	task := repo.Add(title, description)
	fmt.Printf("Task added with ID %d\n", task.ID)
}

func markTaskAsDone(repo *task.Repository, scanner *bufio.Scanner) {
	fmt.Print("Enter task ID: ")
	scanner.Scan()
	idStr := scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}

	task, err := repo.Get(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	task.MarkAsDone()
	if err := repo.Update(task); err != nil {
		fmt.Println("Error updating task:", err)
		return
	}

	fmt.Printf("Task '%s' marked as done\n", task.Title)
}

func viewTaskDetails(repo *task.Repository, scanner *bufio.Scanner) {
	fmt.Print("Enter task ID: ")
	scanner.Scan()
	idStr := scanner.Text()

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID format.")
		return
	}

	task, err := repo.Get(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	status := "Not completed"
	if task.Done {
		status = "Completed"
	}

	fmt.Println("\nTask Details:")
	fmt.Println("=============")
	fmt.Println("ID:", task.ID)
	fmt.Println("Title:", task.Title)
	fmt.Println("Description:", task.Description)
	fmt.Println("Status:", status)
	fmt.Println("Created:", task.CreatedAt.Format("Jan 02, 2006 15:04:05"))
}
