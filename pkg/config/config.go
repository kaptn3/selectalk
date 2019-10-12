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
		Host          string
		Name          string
		Username      string
		Password      string
		Dialect       string
		MigrationsURL string
		Port          int
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
		Host:          viperConfig.Get("Database.Host").(string),
		Port:          viperConfig.Get("Database.Port").(int),
		Name:          viperConfig.Get("Database.Name").(string),
		Username:      viperConfig.Get("Database.Username").(string),
		Password:      viperConfig.Get("Database.Password").(string),
		Dialect:       viperConfig.Get("Database.Dialect").(string),
		MigrationsURL: viperConfig.Get("Database.MigrationsURL").(string),
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
	viperConfig.SetDefault("Database.Host", "localhost")
	viperConfig.SetDefault("Database.Port", 5432)
	viperConfig.SetDefault("Database.Name", "selectel")
	viperConfig.SetDefault("Database.Username", "selectel")
	viperConfig.SetDefault("Database.Password", "selectel")
	viperConfig.SetDefault("Database.Dialect", "postgres")
	viperConfig.SetDefault("Database.MigrationsURL", "file://deployments/migrations/postgres")
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
