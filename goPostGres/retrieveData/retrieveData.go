package retrieveData

import (
	"database/sql"
	"fmt"
	"log"
)

type Student struct {
	id          int
	name        string
	roll_number int
}

func RetrieveRecords(db *sql.DB) {

	// We assign the result to 'rows'
	rowsRs, err := db.Query("SELECT * FROM Students")

	if err != nil {
		return
	}
	defer rowsRs.Close()

	// creates placeholder of the sandbox
	snbs := make([]Student, 0)

	// we loop through the values of rows
	for rowsRs.Next() {
		snb := Student{}
		err := rowsRs.Scan(&snb.id, &snb.name, &snb.roll_number)
		if err != nil {
			log.Println(err)
			return
		}
		snbs = append(snbs, snb)
	}

	for _, snb := range snbs {
		printCall := fmt.Sprintf("%d %s %d\n", snb.id, snb.name, snb.roll_number)
		fmt.Println(printCall)
	}
}
