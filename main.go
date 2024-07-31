package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"moul.io/banner"
	"os"
	"strings"
	"time"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	UserName string    `json:"userName"`
	Password string    `json:"password"`
}

type Task struct {
	Title        string
	Date         time.Time
	Status       bool
	Content      string
	UserID       uuid.UUID
	CategoryName string
}

type Category struct {
	Name   string
	Title  string
	Colour string
}

var (
	authenticatedUser *User
	UserStorage       []User
	TaskStorage       []Task
	CategoryStorage   []Category
	Commands          = []string{"create-task", "create-category", "register", "list-task", "logout", "exit"}
)

func main() {
	fmt.Println("^-^ Welcome to the Todo App ^-^")
	cmd := flag.String("command", "", "Do something")
	flag.Parse()

	for {
		runCommand(*cmd)
		scanner := bufio.NewReader(os.Stdin)
		fmt.Println("Enter the next command:")
		*cmd, _ = scanner.ReadString('\n')
	}
}

func runCommand(cmd string) {
	cmd = strings.Replace(cmd, "\n", "", -1)

	if cmd != "register" && cmd != "exit" && authenticatedUser == nil && cmd != "show-commands" && cmd != "logout" {
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

	case "show-commands":
		showCommands()

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

	fmt.Println("Enter the category name for the task:")
	categoryName, _ := scanner.ReadString('\n')

	for _, category := range CategoryStorage {
		if categoryName == category.Name {
			task := Task{
				Title:        taskTitle,
				Content:      content,
				Date:         time.Now(),
				Status:       true,
				UserID:       authenticatedUser.ID,
				CategoryName: categoryName,
			}

			TaskStorage = append(TaskStorage, task)
		} else {
			fmt.Printf("The category name %s does not exist\n", categoryName)
		}
	}

}

func createCategory() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the category name:")
	categoryName, _ := scanner.ReadString('\n')

	fmt.Println("Enter the category title:")
	categoryTitle, _ := scanner.ReadString('\n')

	fmt.Println("Enter the colour of the category:")
	colour, _ := scanner.ReadString('\n')

	category := Category{
		Name:   categoryName,
		Title:  categoryTitle,
		Colour: colour,
	}

	CategoryStorage = append(CategoryStorage, category)

	fmt.Printf("The category %s is created.\nThe Colour: %s\nThe title: %s", category.Name, category.Colour, categoryTitle)
}

func register() {
	scanner := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your username:")
	username, _ := scanner.ReadString('\n')

	fmt.Println("Enter the password:")
	password, _ := scanner.ReadString('\n')

	user := User{
		ID:       uuid.New(),
		UserName: strings.TrimSpace(username),
		Password: strings.TrimSpace(password),
	}

	data, err := json.Marshal(user)
	data = append(data, '\n')
	if err != nil {
		fmt.Println("couldn't encode user struct in json", err)

		return
	}

	var f *os.File

	//err = os.WriteFile("user.json", data, 0666)
	f, err = os.OpenFile("user.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("could not close the file", err)
		}
	}(f)

	_, err = f.Write(data)
	if err != nil {
		fmt.Println("could not write data to file", err)

		return
	}

	UserStorage = append(UserStorage, user)
}

func login() {
	fmt.Println("You have login first!!!")
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
		} else if len(TaskStorage) == 0 {
			fmt.Println("You have not created tasks yet!")
		}
	}

}

func logout() {
	if authenticatedUser == nil {
		fmt.Println("You are not logged in yet")
		return
	}
	authenticatedUser = nil
}

func showCommands() {
	fmt.Println("==================")
	for _, cmd := range Commands {
		fmt.Println(cmd)
	}
}

func loadUserData(user User) {}
