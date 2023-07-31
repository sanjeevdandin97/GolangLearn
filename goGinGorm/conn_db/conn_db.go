package conn_db

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnDB(host *string, port *int, user *string, password *string, dbname *string) *gorm.DB {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Calcutta", *host, *port, *user, *password, *dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	CheckError(err)

	fmt.Println(`Database Connected!`)

	return db
}

func CloseDB(db *gorm.DB) {
	// close database
	sqlDB, err := db.DB()
	CheckError(err)
	// Close
	sqlDB.Close()

	fmt.Println(`Database closed!`)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
