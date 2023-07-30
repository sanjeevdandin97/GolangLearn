package updateData

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type fn func(error)

func UpdateRecord(db *sql.DB, checkError fn) {
	// update
	updateStmt := `update "students" set "name"=$1, "roll_number"=$2 where "id"=$3`
	_, e := db.Exec(updateStmt, "Rachel", 24, 1)
	checkError(e)
	fmt.Println(`Update Successful!`)
}
