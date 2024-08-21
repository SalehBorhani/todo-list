package filestore

import (
	"encoding/json"
	"fmt"
	"github.com/salehborhani/todo-list/entity"
	"os"
	"strings"
)

type FileStore struct {
	filePath string
}

// constructor

func New(path string) FileStore {
	return FileStore{filePath: path}
}

func (fs FileStore) Save(u entity.User) {
	var f *os.File

	f, err := os.OpenFile(fs.filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
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

	data, err := json.Marshal(u)
	data = append(data, '\n')
	if err != nil {
		fmt.Println("couldn't encode user struct in json", err)

		return
	}

	_, err = f.Write(data)
	if err != nil {
		fmt.Println("could not write data to file", err)

		return
	}

}

func (fs FileStore) Load() []entity.User {
	var uStore []entity.User

	data, err := os.ReadFile(fs.filePath)
	if err != nil {
		return nil
	}

	// split users in text file based-on \n
	userSlice := strings.Split(string(data), "\n")

	for _, u := range userSlice {
		var user entity.User

		if u == "" {
			break
		}
		err = json.Unmarshal([]byte(u), &user)
		if err != nil {
			fmt.Println("could not decode json", err)

			return nil
		}

		uStore = append(uStore, user)
	}
	return uStore
}
