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
mode = "debug"
port = "9000"
http_timeout = 60
cache_expired = 24
cache_purged = 60

[database]
db_driver = "mysql"
db_host = "localhost"
db_port = "3306"
db_name = "your_db_name"
db_username = "root"
db_password = "root"

[pool]
conn_idle = 10
conn_max = 20
conn_lifetime = 60

[jwt]
day_expired = 60
signature_key = "supersecret"

```
- json
```json
{
    "server": {
        "mode": "debug",
        "port": "9000",
        "http_timeout": 60,
        "cache_expired": 24,
        "cache_purged": 60
    },
    "database": {
        "db_driver": "mysql",
        "db_host": "localhost",
        "db_port": "3306",
        "db_name": "your_db_name",
        "db_username": "root",
        "db_password": "root"
    },
    "pool": {
        "conn_idle": 10,
        "conn_max": 20,
        "conn_lifetime": 60
    },
    "jwt": {
        "day_expired": 60,
        "signature_key": "supersecret"
    }
}

```
- yaml
```yaml
server:
  mode: "debug"
  port: "9000"
  http_timeout: 60
  cache_expired: 24
  cache_purged: 60

database:
  db_driver: "mysql"
  db_host: "localhost"
  db_port: "3306"
  db_name: "your_db_name"
  db_username: "root"
  db_password: "root"

pool:
  conn_idle: 10
  conn_max: 20
  conn_lifetime: 60

jwt:
  day_expired: 60
  signature_key: "supersecret"
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

# JSON WEB TOKEN
- create JSON WEB TOKEN:
```go
auth.NewJWTImpl(signatureKey, dayExpiration).GenerateToken(map[string]interface{})
```
- validate JSON WEB TOKEN:
```go
auth.NewJWTImpl(signatureKey, dayExpiration).ValidateToken(JWToken)
```
- parse JSON WEB Token
```go
auth.NewJWTImpl(signatureKey, dayExpiratio).ParseToken(JWToken)
```
example:
```go
	// create jwt
	dayExpired := 7 // one week
	dataStoreToToken := map[string]interface{}{
		"user":         "John Doe",
		"authenticate": true,
	}
	jwToken, err := NewJWTImpl("superSecret", dayExpired).GenerateToken(dataStoreToToken)
	if err != nil {
		panic(err)
	}
	fmt.Println(jwToken)

	// validate jwt
	valid, _ := NewJWTImpl("superSecret", dayExpired).ValidateToken(jwToken)
	if !valid {
		fmt.Println("token not valid")
	}
	fmt.Println("token valid")

	// parse jwt
	planData, _ := NewJWTImpl("superSecret", dayExpired).ParseToken(jwToken)
	fmt.Println(planData)
```
# Caching
- set key and value cache
```go
	simpleCache := simplecache.NewSimpleCache(simplecache.SimpleCache{
		ExpiredAt: 10, // How long will the data remain stored in the cache before being deleted
		PurgeTime: 30, // How long will the data remain stored in the cache before it is deleted
	})

	cache := simpleCache.Open()
	simpleCache.Set("hello", "world")
```

- get the value
```go
  value := simpleCache.Get("hello")
  fmt.Println(vaue)
```

- delete key and value
```go
	simpleCache.Delete("hello")
```


# Example
 [example](https://github.com/zakirkun/infra-go/tree/main/example)
---
# How to run ?
```sh
go run main.go -c config.toml
```