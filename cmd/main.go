package main

import (
	"fmt"

	handler "github.com/mycontroller-org/mycontroller/cmd/app"
	storage "github.com/mycontroller-org/mycontroller/pkg/storage"
)



func main() {
	fmt.Println("Welcome to MyController 2.x :)")
	config := map[string]string{
		"database": "mydb",
		"uri":      "mongodb+srv://testuser:testuser@cluster0-sk7af.mongodb.net/test?retryWrites=true&w=majority",
	}
	err := storage.GetClient(config)
	if err != nil {
		panic(err)
	}

	handler.StartHandler()
}
