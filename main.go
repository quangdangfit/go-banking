package main

import (
	"go-banking/api"
	"go-banking/database"
)

func main() {
	//migrations.Migrate()
	//migrations.MigrateTransactions()
	database.InitDatabase()
	api.StartApi()
}
