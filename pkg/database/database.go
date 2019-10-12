// Package orm aims to work with relational databases via gorm project.
package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rodsher/selectel/pkg/config"
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
	migrate()
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

func migrate() {
	db.AutoMigrate()
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