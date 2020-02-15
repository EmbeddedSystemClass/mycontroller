package main

import (
	"errors"
	"flag"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"

	"github.com/mycontroller-org/mycontroller/cmd/app/handler"
	"github.com/mycontroller-org/mycontroller/pkg/metrics"
	"github.com/mycontroller-org/mycontroller/pkg/storage"
)

// Database to be used
type Database struct {
	Storage string `yaml:"storage"`
	Metrics string `yaml:"metrics"`
}

// Config of the system
type Config struct {
	Web       handler.WebConfig   `yaml:"web"`
	Database  Database            `yaml:"database"`
	Databases []map[string]string `yaml:"databases"`
	Logger    map[string]string   `yaml:"logger"`
}

var cfg Config

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
		zap.L().Fatal("Configuration file not supplied")
	}
	zap.L().Debug("Configuration file path:", zap.String("file", *cf))
	d, err := ioutil.ReadFile(*cf)
	if err != nil {
		zap.L().Fatal("Error on reading configuration file", zap.Error(err))
	}

	err = yaml.Unmarshal(d, &cfg)
	if err != nil {
		zap.L().Fatal("Failed to unmarshal yaml data", zap.Error(err))
	}

	// Get storage and metric database config
	sCfg, err := getDatabaseConfig(cfg.Database.Storage, &cfg)
	if err != nil {
		zap.L().Fatal("Problem with storage database config", zap.String("name", cfg.Database.Storage), zap.Error(err))
	}
	mCfg, err := getDatabaseConfig(cfg.Database.Metrics, &cfg)
	if err != nil {
		zap.L().Fatal("Problem with metrics database config", zap.String("name", cfg.Database.Metrics), zap.Error(err))
	}

	// Init storage database
	err = storage.Init(sCfg)
	if err != nil {
		zap.L().Fatal("Error on storage db init", zap.Error(err))
	}

	// Init metrics database
	err = metrics.Init(mCfg)
	if err != nil {
		zap.L().Fatal("Error on metrics db init", zap.Error(err))
	}
}

func main() {
	defer zap.L().Sync()
	err := handler.StartHandler(&cfg.Web)
	if err != nil {
		zap.L().Fatal("Error on starting http handler", zap.Error(err))
	}
}

func getDatabaseConfig(name string, cfg *Config) (map[string]string, error) {
	for _, d := range cfg.Databases {
		if d["name"] == name {
			return d, nil
		}
	}
	return nil, errors.New("Config not found")
}
