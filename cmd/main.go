package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/mycontroller-org/mycontroller/cmd/app/handler"
	"github.com/mycontroller-org/mycontroller/pkg/storage"
)

// WebConfig input
type WebConfig struct {
	BindAddress  string `yaml:"bindAddress"`
	Port         uint   `yaml:"port"`
	WebDirectory string `yaml:"webDirectory"`
}

// Config of the system
type Config struct {
	Database map[string]string `yaml:"database"`
	Web      WebConfig         `yaml:"web"`
}

func main() {
	fmt.Println("Welcome to MyController 2.x :)")
	cf := flag.String("config", "./config.yaml", "Configuration file")
	flag.Parse()
	if cf == nil {
		panic(errors.New("Configuration file not supplied"))
	}
	fmt.Println(flag.Args())
	fmt.Println("File Name:", *cf)
	d, err := ioutil.ReadFile(*cf)
	if err != nil {
		panic(err)
	}
	var config Config
	yaml.Unmarshal(d, &config)
	db := map[string]string{
		"database": config.Database["name"],
		"uri":      config.Database["uri"],
	}
	err = storage.Init(db)
	if err != nil {
		panic(err)
	}

	handler.StartHandler()
}
