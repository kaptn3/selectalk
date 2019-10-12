// Package orm aims to work with relational databases via gorm project.
package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rodsher/selectel/pkg/config"
	"github.com/rodsher/selectel/pkg/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// GetInstance returns pointer to an instance of gorm.DB structure.
func GetInstance() *gorm.DB {
	return db
}

// db is a private variable of package orm. Must be used across packages
// via package level method GetDB.
var db *gorm.DB

// Init function opens connection with a DBMS, setups specific database options and
// runs migration.
func Init() {
	open()
	setup()
	autoMigrate()
	seed()
}

// open function builds connection string for a database and tries to open connection
// with a database server. In case when connection established it initializes package level variable "db".
// Otherwise, the program exited with non-zero code.
func open() {
	conf := config.GetInstance().Database
	connStr := connStrFromConf(conf)

	database, err := gorm.Open(conf.Dialect, connStr)
	if err != nil {
		log.Fatal("failed connect to database", err)
	}

	db = database
}

func autoMigrate() {
	db.AutoMigrate(
		&model.User{},
		&model.TaskPriority{},
		&model.TaskStatus{},
		&model.Task{},
		&model.Achievement{},
		&model.Course{},
	)
}

func seed() {
	conf := config.GetInstance().Database
	url := concatDBMigrationURL()
	m, err := migrate.New(conf.MigrationsURL, url)
	if err != nil {
		log.Fatal(err)
	}

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			log.Println("Seeders up to date")
			return
		}

		log.Fatal(err)
	}
}

func setup() {
	db.SingularTable(true)
	db.LogMode(true)
}

func connStrFromConf(c config.DatabaseConfig) string {
	port := strconv.Itoa(int(c.Port))
	connStr := "host=%s port=%s user=%s dbname=%s password=%s"
	return fmt.Sprintf(connStr, c.Host, port, c.Username, c.Name, c.Password)
}

func concatDBMigrationURL() string {
	conf := config.GetInstance().Database
	port := strconv.Itoa(conf.Port)
	url := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	return fmt.Sprintf(url, conf.Username, conf.Password, conf.Host, port, conf.Name)
}
