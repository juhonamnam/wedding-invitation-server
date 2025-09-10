package sqldb

import (
	"database/sql"
)

var (
	sqlDb *sql.DB
)

func SetDb(db *sql.DB) {
	sqlDb = db
	err := initializeGuestbookTable()
	if err != nil {
		panic(err)
	}
	err = initializeAttendanceTable()
	if err != nil {
		panic(err)
	}
}

func GetDb() *sql.DB {
	return sqlDb
}
