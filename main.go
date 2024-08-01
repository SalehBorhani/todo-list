package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
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

const DATABASE = "user.txt"

func main() {
	// Load user data, if we have a database of users
	UserStorage = loadUsers(DATABASE)

	// Start the main program
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
		}
	}
	fmt.Printf("The category name %s does not exist.\nPlease create one!\n", strings.TrimSpace(categoryName))

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

	fmt.Printf("The category %s is created.\nThe Colour: %s\nThe title: %s", strings.TrimSpace(category.Name), category.Colour, categoryTitle)
}

func register() {
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

	user := User{
		ID:       uuid.New(),
		UserName: username,
		Password: hash,
	}

	data, err := json.Marshal(user)
	data = append(data, '\n')
	if err != nil {
		fmt.Println("couldn't encode user struct in json", err)

		return
	}

	writeData(data)
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

func loadUsers(file string) []User {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	// split users in text file based-on \n
	userSlice := strings.Split(string(data), "\n")

	for _, u := range userSlice {
		var user User

		if u == "" {
			break
		}
		err = json.Unmarshal([]byte(u), &user)
		if err != nil {
			fmt.Println("could not decode json", err)

			return nil
		}

		UserStorage = append(UserStorage, user)
	}
	return UserStorage
}

func writeData(data []byte) {
	var f *os.File

	f, err := os.OpenFile(DATABASE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("could not open file", err)

		return
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("could not close the file", err)

			return
		}
	}(f)

	_, err = f.Write(data)
	if err != nil {
		fmt.Println("could not write data to file", err)

		return
	}

}

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password))

	return hex.EncodeToString(hash[:])
}
