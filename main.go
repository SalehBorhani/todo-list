package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/salehborhani/todo-list/contract"
	"github.com/salehborhani/todo-list/entity"
	"github.com/salehborhani/todo-list/repository/filestorage"
	"moul.io/banner"
	"os"
	"strings"
	"time"
)

var (
	authenticatedUser *entity.User
	UserStorage       []entity.User
	TaskStorage       []entity.Task
	CategoryStorage   []entity.Category
	Commands          = []string{"create-task", "create-category", "register", "list-task", "logout", "exit"}
)

const DATABASE = "user.txt"

func main() {
	// Start the main program
	fmt.Println("^-^ Welcome to the Todo App ^-^")
	cmd := flag.String("command", "", "Do something")
	flag.Parse()

	// Load user data, if we have a database of users
	var fileStore = filestorage.New(DATABASE)
	users := fileStore.Load()
	UserStorage = append(UserStorage, users...)

	for {
		runCommand(fileStore, *cmd)
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Enter the next command:")
		scanner.Scan()
		*cmd = scanner.Text()
	}
}

func runCommand(store contract.UserWriteStore, cmd string) {

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
		register(store)

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
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the task title:")
	scanner.Scan()
	taskTitle := scanner.Text()

	fmt.Println("Enter the content of the task:")
	scanner.Scan()
	content := scanner.Text()

	fmt.Println("Enter the category name for the task:")
	scanner.Scan()
	categoryName := scanner.Text()

	for _, category := range CategoryStorage {
		if categoryName != category.Name {
			fmt.Printf("The category name %s does not exist.\nPlease create one!\n", strings.TrimSpace(categoryName))
		}
	}

	task := entity.Task{
		Title:        taskTitle,
		Content:      content,
		Date:         time.Now(),
		Status:       true,
		UserID:       authenticatedUser.ID,
		CategoryName: categoryName,
	}
	TaskStorage = append(TaskStorage, task)
}

func createCategory() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the category name:")
	scanner.Scan()
	categoryName := scanner.Text()

	fmt.Println("Enter the category title:")
	scanner.Scan()
	categoryTitle := scanner.Text()

	fmt.Println("Enter the colour of the category:")
	scanner.Scan()
	colour := scanner.Text()

	category := entity.Category{
		Name:   categoryName,
		Title:  categoryTitle,
		Colour: colour,
	}

	CategoryStorage = append(CategoryStorage, category)

	for _, c := range CategoryStorage {
		fmt.Println(c.Name)
	}
}

func register(store contract.UserWriteStore) {
	fmt.Println(banner.Inline("register"))
	scanner := bufio.NewScanner(os.Stdin)
	var username, password string

	fmt.Println("please enter username:")
	scanner.Scan()
	username = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	hash := HashPassword(password)

	user := entity.User{
		ID:       uuid.New(),
		UserName: username,
		Password: hash,
	}

	store.Save(user)
	UserStorage = append(UserStorage, user)
}

func login() {
	fmt.Println("==========================")
	fmt.Println(banner.Inline("login"))
	scanner := bufio.NewScanner(os.Stdin)
	var username, password string

	fmt.Println("please enter username:")
	scanner.Scan()
	username = scanner.Text()

	fmt.Println("please enter the password")
	scanner.Scan()
	password = scanner.Text()

	hash := HashPassword(password)

	for _, user := range UserStorage {
		if user.UserName == username && user.Password == hash {
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

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}
