package config

import (
	"strconv"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		Env      string
		Database DatabaseConfig
	}

	DatabaseConfig struct {
		Host, Name, Username, Password, Dialect string
		Port                                    int
	}
)

var (
	viperConfig = viper.New()
	config      *Config
)

func Init() {
	setDefault()
	setConfig()
}

func GetInstance() *Config {
	return config
}

func setDefault() {
	setDefaultEnvironment()
	setDefaultDatabase()
}

func setConfig() {
	env := viperConfig.Get("Environment").(string)
	database := DatabaseConfig{
		Dialect:  viperConfig.Get("Database.Dialect").(string),
		Host:     viperConfig.Get("Database.Host").(string),
		Port:     viperConfig.Get("Database.Port").(int),
		Name:     viperConfig.Get("Database.Name").(string),
		Username: viperConfig.Get("Database.Username").(string),
		Password: viperConfig.Get("Database.Password").(string),
	}

	config = &Config{
		Env:      env,
		Database: database,
	}
}

func setDefaultEnvironment() {
	viperConfig.SetDefault("Environment", "production")
}

func setDefaultDatabase() {
	viperConfig.SetDefault("Database.Dialect", "postgres")
	viperConfig.SetDefault("Database.Host", "localhost")
	viperConfig.SetDefault("Database.Port", 5432)
	viperConfig.SetDefault("Database.Name", "selectel")
	viperConfig.SetDefault("Database.Username", "selectel")
	viperConfig.SetDefault("Database.Password", "selectel")
}

// convertDatabasePortToInteger inspects type of Metric.Port value and if it is string converts
// string to an integer.
func convertDatabasePortToInteger() {
	if v, ok := viperConfig.Get("Database.Port").(string); ok {
		port, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}

		viperConfig.Set("Database.Port", port)
	}
}
