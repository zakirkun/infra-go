# Insfrastructure SDK for Golang Project 
## Installation
```sh
go get github.com/zakirkun/infra-go@latest
```
---
Create your configuration
- toml
```toml
[server]
mode="debug"
port="8000"
http_timeout=60

[database]
db_driver="db_drivers"
db_host="localhost"
db_port="3306"
db_name="your_db"
db_username="db_user"
db_password="db_password"

[pool]
conn_idle=10
conn_max=20
conn_lifetime=60
```
- json
```json
{
  "server": {
    "mode": "debug",
    "port": "9000",
    "http_timeout": 60
  },
  "database": {
    "db_driver": "db_driver",
    "db_host": "localhost",
    "db_port": "3306",
    "db_name": "your_db",
    "db_username": "db_user",
    "db_password": "db_password"
  },
  "pool": {
    "conn_idle": 25,
    "conn_max": 50,
    "conn_lifetime": 60
  }
}

```
- yaml
```yaml
server:
  mode: "debug"
  port: "9000"
  http_timeout: 60
database:
  db_driver: "db_driver"
  db_host: "localhost"
  db_port: "3306"
  db_name: "your_db"
  db_username: "db_user"
  db_password: "db_password"
pool:
  conn_idle: 25
  conn_max: 50
  conn_lifetime: 60
```
- db_drivers : \
 We currently use 2 database drivers, `mysql` and `postgres`

 ### What is Pool Connection ?
 ---
 Connection pooling is a mechanism used in applications to manage connections to data sources like databases. It stores ready-to-use connections, enabling the application to quickly acquire and release connections to the data source.
---

- pool \
conn_idle : 10 (maximum number of connections in the idle connection pool)\
conn_max : 10 (maximum number of open connections to the database)\
conn_lifetime : 60 (maximum time in minutes for reusing a connection)

# Example
#### your app structure base
```sh
├── config.toml
├── go.mod
├── go.sum
├── main.go
├── models
│   └── users.go
└── routers
    └── routers.go
```

`main.go`
```go
package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/zakirkun/infra-go/infrastructure"
	"github.com/zakirkun/infra-go/pkg/config"
	"github.com/zakirkun/infra-go/pkg/database"
	"github.com/zakirkun/infra-go/pkg/server"

    "yourApp/models"
	"yourApp/routers"

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
	err := database.NewMigration(&models.Users{})
	log.Printf("Migration Info : %v", err)
}

func SetDatabase() database.DBModel {
	return database.DBModel{
		ServerMode:   config.GetString("server.mode"),
		Driver:       config.GetString("database.db_driver"),
		Host:         config.GetString("database.db_host"),
		Port:         config.GetString("database.db_port"),
		Name:         config.GetString("database.db_name"),
		Username:     config.GetString("database.db_username"),
		Password:     config.GetString("database.db_password"),
		MaxIdleConn:  config.GetInt("pool.conn_idle"),
		MaxOpenConn:  config.GetInt("pool.conn_max"),
		ConnLifeTime: config.GetInt("pool.conn_lifetime"),
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
```
---
`models.go`
```go
package models

import "gorm.io/gorm"

type Users struct {
	*gorm.Model
	Username string `gorm:"not null;size:255;default:anonymous"`
	Email    string `gorm:"not null;unique"`
}
```
---
`routers.go`\
for example we use [echo](github.com/labstack/echo) web framework
```go
package routers

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRouters() http.Handler {
	e := echo.New()

	// middleware section
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogMethod:    true,
		LogUserAgent: true,
		LogRoutePath: true,
		LogHost:      true,
		BeforeNextFunc: func(c echo.Context) {
			if c.Request().Header.Get(echo.HeaderXRequestID) == "" {
				c.Request().Header.Set(echo.HeaderXRequestID, uuid.NewString())
			}
		},
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			fmt.Printf("[%v] REQUEST: uri: %v, Host: %v, Method: %v, UserAgent: %v, RoutePath: %v, IP: %v\n", v.RequestID, v.URI, v.Host, v.Method, v.UserAgent, v.RoutePath, v.RemoteIP)
			return nil
		},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"messages": "Hello World!"})
	})

	return e
}
```
---
# How to run ?
```sh
go run main.go -c config.toml
```