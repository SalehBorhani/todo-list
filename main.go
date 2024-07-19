package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"moul.io/banner"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       uuid.UUID
	UserName string
	Password string
}

type Task struct {
	Title   string
	Date    time.Time
	Status  bool
	Content string
	UserID  uuid.UUID
}

type Category struct {
	Title  string
	Tasks  []Task
	Colour string
}

var authenticatedUser *User

var UserStorage []User
var TaskStorage []Task

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

	if cmd != "register" && cmd != "exit" && authenticatedUser == nil {
		login()
		if authenticatedUser == nil {
			return
		}
	}

	switch cmd {
	case "create-task":
		createTask()

	case "create-category":
		createCategory()

	case "register":
		register()

	case "list-task":
		listTask()

	case "logout":
		logout()

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
		Date:    time.Now(),
		Status:  true,
		UserID:  authenticatedUser.ID,
	}

	TaskStorage = append(TaskStorage, task)

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

	UserStorage = append(UserStorage, user)

	fmt.Printf("ID: %s, Username: %s, Password: %s\n", user.ID, user.UserName, user.Password)
}

func login() {
	fmt.Println("==========================")
	fmt.Println(banner.Inline("login"))
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your username:")
	username, _ := scanner.ReadString('\n')

	fmt.Println("Enter the password:")
	password, _ := scanner.ReadString('\n')

	for _, user := range UserStorage {
		if user.UserName == username && user.Password == password {
			fmt.Println("You are login")
			authenticatedUser = &user
			break
		}
	}
	if authenticatedUser == nil {
		fmt.Println("Invalid provided credentials")
	}
}

func listTask() {
	for _, task := range TaskStorage {
		if authenticatedUser.ID == task.UserID {
			fmt.Printf("The task %s is created at %s by %s.\nThe Status: %t\nThe Content: %s", task.Title, task.Date, task.UserID, task.Status, task.Content)
		}
	}
}

func logout() {
	authenticatedUser = nil
}
