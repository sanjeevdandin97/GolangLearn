package deleteData

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type fn func(error)

func DeleteData(db *sql.DB, checkError fn) {
	// Insert hardcoded
	deleteStmt := `DELETE FROM students`
	_, e := db.Exec(deleteStmt)
	checkError(e)
	fmt.Println(`Delete Successful!`)
}
