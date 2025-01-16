package config

import (
	"log"

	"github.com/umed-hotamov/house-rental/internal/app/server"

	"github.com/spf13/viper"
)

type Config struct {
  Server server.Config
}

func GetConfig() *Config {
  if err := bindEnvConfig(); err != nil {
    log.Fatalf("error reading config: %v", err)
  }

  cfg := new(Config)
  if err := viper.Unmarshal(cfg); err != nil {
    log.Fatalf("error unmarshaling config: %v", err)
  }

  return cfg
}

func bindEnvConfig() error {
  bindings := make(map[string]string)

  bindings["server.host"] = "HOST"
  bindings["server.port"] = "PORT"

  for name, binding := range bindings {
    if err := viper.BindEnv(name, binding); err != nil {
      return err
    }
  }

  return nil 
}
