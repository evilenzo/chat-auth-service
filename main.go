package main

import (
	"main/db_operator"
	"main/server"

	psql "github.com/evilenzo/psql-connector"
)

func main() {
	postgres := psql.ConnectFromEnv()
	dbo := db_operator.CreateDatabaseOperator(postgres)
	httpServer := server.CreateServer(dbo)
	httpServer.Run()
}
