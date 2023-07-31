package insert

import "gorm.io/gorm"

type Students struct {
	Name       string
	RollNumber int
}

func InsertRecord(db *gorm.DB) string {
	students := []Students{
		{Name: "Rachel", RollNumber: 222},
		{Name: "Ras Al Ghul", RollNumber: 223},
	}

	db.Create(students) // pass a slice to insert multiple row

	return `Successful`

	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count
}
