# Insfrastructure SDK for Golang Project 
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

[pool]
conn_idle=10
conn_max=20
conn_lifetime=60
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