package main

import (
	"errors"
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"

	"github.com/mycontroller-org/mycontroller/cmd/app/handler"
	"github.com/mycontroller-org/mycontroller/pkg/storage"
)

// Config of the system
type Config struct {
	Database map[string]string `yaml:"database"`
	Web      handler.WebConfig `yaml:"web"`
}

var config Config

func init() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)

	zap.L().Info("Welcome to MyController 2.x :)")

	cf := flag.String("config", "./config.yaml", "Configuration file")
	flag.Parse()
	if cf == nil {
		panic(errors.New("Configuration file not supplied"))
	}
	zap.L().Debug("Configuration file path:", zap.String("file", *cf))
	d, err := ioutil.ReadFile(*cf)
	if err != nil {
		panic(err)
	}

	yaml.Unmarshal(d, &config)
	db := map[string]string{
		"database": config.Database["name"],
		"uri":      config.Database["uri"],
	}
	err = storage.Init(db)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer zap.L().Sync()
	handler.StartHandler(&config.Web)
}
