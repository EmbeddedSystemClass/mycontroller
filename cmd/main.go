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
		"uri":      "test",
	}
	err := storage.GetClient(config)
	if err != nil {
		panic(err)
	}

	handler.StartHandler()
}
