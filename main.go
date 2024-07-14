package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
)

type User struct {
	ID       uuid.UUID
	UserName string
	Password string
}

type Task struct {
	Title   string
	Status  bool
	Content string
}

type Category struct {
	Title  string
	Tasks  []Task
	Colour string
}

var StorageUsers []User

func main() {
	fmt.Println("^-^ Welcome to the Todo App ^-^")
	cmd := flag.String("command", "", "Do something")
	flag.Parse()

	for {
		runCommand(*cmd)
		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the command:")
		*cmd, _ = scanner.ReadString('\n')
	}
}

func runCommand(cmd string) {
	cmd = strings.Replace(cmd, "\n", "", -1)

	switch cmd {
	case "create-task":
		createTask()

	case "create-category":
		createCategory()

	case "register":
		register()

	case "login":
		login()

	case "exit":
		os.Exit(0)

	default:
		fmt.Printf("The command is invalid: %s\n", cmd)
	}
}

func createTask() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the task title:")
	taskTitle, _ := scanner.ReadString('\n')

	fmt.Println("Enter the content of the task:")
	content, _ := scanner.ReadString('\n')

	task := Task{
		Title:   taskTitle,
		Content: content,
		Status:  true,
	}

	fmt.Printf("The task %s is created.\nThe Status: %t\nThe Content: %s", task.Title, task.Status, task.Content)

}

func createCategory() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the category title:")
	categoryTitle, _ := scanner.ReadString('\n')

	fmt.Println("Enter the colour of the category:")
	colour, _ := scanner.ReadString('\n')

	category := Category{
		Title:  categoryTitle,
		Colour: colour,
	}

	fmt.Printf("The category %s is created.\nThe Colour: %s", category.Title, category.Colour)
}

func register() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your username:")
	username, _ := scanner.ReadString('\n')

	fmt.Println("Enter the password:")
	password, _ := scanner.ReadString('\n')

	user := User{
		ID:       uuid.New(),
		UserName: username,
		Password: password,
	}

	StorageUsers = append(StorageUsers, user)

	fmt.Printf("ID: %s, Username: %s, Password: %s\n", user.ID, user.UserName, user.Password)
}

func login() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your username:")
	username, _ := scanner.ReadString('\n')

	fmt.Println("Enter the password:")
	password, _ := scanner.ReadString('\n')

	for user := range StorageUsers {
		if StorageUsers[user].UserName != username {
			fmt.Println("You are not registered yet!. login first ^=^")
			break
		}
		if StorageUsers[user].UserName == username && StorageUsers[user].Password == password {
			fmt.Println("You are login")
			break
		} else {
			fmt.Println("invalid provided credentials")
			break
		}
	}
}
