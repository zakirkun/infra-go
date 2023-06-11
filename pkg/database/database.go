package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBModel struct {
	ServerMode   string
	Driver       string
	Host         string
	Port         string
	Name         string
	Username     string
	Password     string
	MaxIdleConn  int
	MaxOpenConn  int
	ConnLifeTime int
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

	conPool, err := db.DB()
	if err != nil {
		log.Fatalf("Cannot Create Connection Pool to DB With Message %s", err.Error())
		return nil, &err
	}

	/** SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	**/
	conPool.SetMaxIdleConns(c.MaxIdleConn)

	/** SetMaxOpenConns sets the maximum number of open connections to the database.
	**/
	conPool.SetMaxOpenConns(c.MaxOpenConn)

	/** SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	**/
	conPool.SetConnMaxLifetime(time.Duration(c.ConnLifeTime) * time.Minute)

	return db, nil
}
