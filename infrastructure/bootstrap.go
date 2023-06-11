package infrastructure

import (
	"log"
	"os"

	"github.com/zakirkun/infra-go/example/app"
	simplecache "github.com/zakirkun/infra-go/pkg/cache/simple_cache"
	"github.com/zakirkun/infra-go/pkg/database"
	"github.com/zakirkun/infra-go/pkg/server"
)

type Infrastructure interface {
	Database()
	SimpleCache()
	WebServer()
}

type infrastructureContext struct {
	database database.DBModel
	cache    simplecache.SimpleCache
	server   server.ServerContext
}

func NewInfrastructure(database database.DBModel,
	cache simplecache.SimpleCache,
	server server.ServerContext,
) Infrastructure {
	return infrastructureContext{
		database: database,
		cache:    cache,
		server:   server,
	}
}

func (i infrastructureContext) Database() {
	_, err := i.database.OpenDB()
	if err != nil {
		os.Exit(1)
	}

	database.DB = &i.database

}

func (i infrastructureContext) SimpleCache() {
	simplecache.Cache = i.cache.Open()

	if simplecache.Cache == nil {
		log.Fatal("Failed create cache")
	}

	app.SimpleCache = i.cache
}

func (i infrastructureContext) WebServer() {
	i.server.Run()
}
