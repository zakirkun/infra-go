# Insfrastructure Boilerplate for Golang Project 
Set your configuration in config.toml
``` example
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
```
- db_drivers : \
 We currently use 2 database drivers, `mysql` and `postgres`

## Structure 
```
├── cmd
│   └── main.go
├── config.toml
├── go.mod
├── go.sum
├── infrastructure
│   └── bootstrap.go
├── internal
│   ├── config
│   │   └── toml.go
│   ├── database
│   │   ├── database.go
│   │   ├── gen.go
│   │   ├── globals.go
│   │   └── migration.go
│   └── server
│       └── server.go
└── routers
    └── routers.go
```
# Explanation for each layer in the Infra-Go:
- ### cmd
The `main.go` file in this directory serves as the entry point of the application. It initializes and executes the main components of the project. 

- ### infrastructure
Responsible for initializing or preparing the application environment. This includes database initialization, network connection settings, or other necessary configurations when starting the application.
- ### Internal
Contains the sub-layers `config`, `database`, and `server`

- config
    
     Reads the configuration file.
- database
    
    Contains files related to database operations
- server

    Runs an HTTP server or other required communication protocols for the application

- ### Routers
    Configuration of routes or HTTP endpoints