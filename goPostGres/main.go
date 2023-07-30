package main

import (
	"goPostGres/configs"
	"goPostGres/connDB"
	"goPostGres/deleteData"
	"goPostGres/insertData"
	"goPostGres/retrieveData"
	"goPostGres/updateData"
)

func main() {
	dbconfig := configs.DBConfig

	db := connDB.ConnDB(&dbconfig.Host, &dbconfig.Port, &dbconfig.User, &dbconfig.Password, &dbconfig.DBName)

	insertData.InsertData(db, connDB.CheckError)

	retrieveData.RetrieveRecords(db)

	updateData.UpdateRecord(db, connDB.CheckError)

	retrieveData.RetrieveRecords(db)

	deleteData.DeleteData(db, connDB.CheckError)

	connDB.CloseDB(db)
}
