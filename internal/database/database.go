package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBModel struct {
	ServerMode string
	Driver     string
	Host       string
	Port       string
	Name       string
	Username   string
	Password   string
}

func (c *DBModel) OpenDB() (*gorm.DB, *error) {

	var connection gorm.Dialector

	switch c.Driver {
	case "postgres":
		connectionUrl := fmt.Sprintf(POSGRES_CONFIG, c.Username, c.Password, c.Name, c.Host, c.Port, "disable")
		connection = postgres.Open(connectionUrl)
	case "mysql":
		connectionUrl := fmt.Sprintf(MYSQL_CONFIG, c.Username, c.Password, c.Host, c.Port, c.Name)
		connection = mysql.Open(connectionUrl)
	default:
		log.Fatal("No Database Selected!, Please check config.toml")
		os.Exit(1)
	}

	db, err := gorm.Open(connection, &gorm.Config{})
	if err != nil {
		log.Fatalf("Cannot Connect to DB With Message %s", err.Error())
		return nil, &err
	}

	return db, nil
}
