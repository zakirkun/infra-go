package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/zakirkun/infra-go/example/models"
	routers "github.com/zakirkun/infra-go/example/routers"
	"github.com/zakirkun/infra-go/infrastructure"
	"github.com/zakirkun/infra-go/internal/config"
	"github.com/zakirkun/infra-go/internal/database"
	"github.com/zakirkun/infra-go/internal/server"
)

var configFile *string

func init() {
	configFile = flag.String("c", "config.toml", "configuration file")
	flag.Parse()
}

func main() {
	setConfig()

	infra := infrastructure.NewInfrastructure(SetDatabase(), SetWebServer())
	infra.Database()
	SetMigration()

	infra.WebServer()
}

func setConfig() {
	cfg := config.NewConfig(*configFile)
	if err := cfg.Initialize(); err != nil {
		log.Fatalf("Error reading config : %v", err)
		os.Exit(1)
	}
}

func SetMigration() {
	err := database.NewMigration(models.Account{})
	log.Printf("Migration Info : %v", err)
}

func SetDatabase() database.DBModel {
	return database.DBModel{
		ServerMode: config.GetString("server.mode"),
		Driver:     config.GetString("database.db_driver"),
		Host:       config.GetString("database.db_host"),
		Port:       config.GetString("database.db_port"),
		Name:       config.GetString("database.db_name"),
		Username:   config.GetString("database.db_username"),
		Password:   config.GetString("database.db_password"),
	}
}

func SetWebServer() server.ServerContext {
	return server.ServerContext{
		Host:         ":" + config.GetString("server.port"),
		Handler:      routers.InitRouters(),
		ReadTimeout:  time.Duration(config.GetInt("server.http_timeout")),
		WriteTimeout: time.Duration(config.GetInt("server.http_timeout")),
	}
}
