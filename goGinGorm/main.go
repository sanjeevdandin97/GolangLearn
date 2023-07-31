package main

import (
	"goGinGorm/configs"
	"goGinGorm/conn_db"
	"goGinGorm/server"
)

func main() {
	dbConfig := configs.DBConfig

	db := conn_db.ConnDB(
		&dbConfig.Host,
		&dbConfig.Port,
		&dbConfig.User,
		&dbConfig.Password,
		&dbConfig.DBName,
	)

	server.RouterInit(configs.GoRoutes, db)

	conn_db.CloseDB(db)
}
