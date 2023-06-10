package infrastructure

import (
	"os"

	"github.com/zakirkun/infra-go/internal/database"
	"github.com/zakirkun/infra-go/internal/server"
)

type Infrastructure interface {
	Database()
	WebServer()
}

type infrastructureContext struct {
	database database.DBModel
	server   server.ServerContext
}

func NewInfrastructure(database database.DBModel,
	server server.ServerContext,
) Infrastructure {
	return infrastructureContext{
		database: database,
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

func (i infrastructureContext) WebServer() {
	i.server.Run()
}
