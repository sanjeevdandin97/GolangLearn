package insertData

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type fn func(error)

func InsertData(db *sql.DB, checkError fn) {
	// Insert hardcoded
	insertStmt := `insert into "students"("name", "roll_number") values('Jacob', 20)`
	_, e := db.Exec(insertStmt)
	checkError(e)
	fmt.Println(`Insert Successful!`)

	// Insert dynamic
	insertDynStmt := `insert into "students"("name", "roll_number") values($1, $2)`
	_, e = db.Exec(insertDynStmt, "Jack", 21)
	checkError(e)
	fmt.Println(`Insert Successful!`)
}
