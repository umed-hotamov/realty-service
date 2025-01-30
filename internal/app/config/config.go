package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
  Server   ServerConfig
  Postgres PosrtgresConfig
  Logger   LoggerConfig
}

type ServerConfig struct {
  Host string
  Port string
}

type PosrtgresConfig struct {
  Host     string
  Port     string
  DBName   string
  User     string
  Password string
}

type LoggerConfig struct {
  Path     string
  FileName string
  Level    string
}

func GetConfig() *Config {
  viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.SetConfigType("yml")

  viper.AutomaticEnv()

  if err := bindEnvConfig(); err != nil {
    log.Fatalf("error reading config: %v", err)
  }

  log.Println("reading config file: config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading config file: %v", err)
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
