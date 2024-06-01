package sqldb

import (
	"fmt"
	"time"
)

func initializeAttendanceTable() error {
	_, err := sqlDb.Exec(`
		CREATE TABLE IF NOT EXISTS attendance (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			side VARCHAR(10),
			name VARCHER(20),
			meal VARCHAR(20),
			count INTEGER,
			timestamp INTEGER
		)
	`)
	return err
}

func CreateAttendance(side, name, meal string, count int) error {
	_, err := sqlDb.Exec(`
		INSERT INTO attendance (side, name, meal, count, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`, side, name, meal, count, time.Now().Unix())
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
