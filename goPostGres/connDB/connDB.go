package connDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnDB(host *string, port *int, user *string, password *string, dbname *string) *sql.DB {
	fmt.Println(*host, *port, *user, *password, *dbname)
	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	return db
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CloseDB(db *sql.DB) {
	// close database
	defer db.Close()

	fmt.Println(`Database closed!`)
}
