package db

import (
	"database/sql"
	"fmt"

	_"modernc.org/sqlite"

)

var DB *sql.DB
 
func InitDB() {
    var err error
    DB, err = sql.Open("sqlite", "api.db")
 
    if err != nil {
        panic("Could not connect to database.")
    }
 
    createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create events table")
	}
}