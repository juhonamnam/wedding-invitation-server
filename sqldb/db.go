package sqldb

import (
	"database/sql"
	"github.com/juhonamnam/wedding-invitation-server/env"
)

var (
	sqlDb *sql.DB
)

func SetDb(db *sql.DB) {
	sqlDb = db
	if env.UseGuestbook {
		err := initializeGuestbookTable()
		if err != nil {
			panic(err)
		}
	}
	if env.UseAttendance {
		err := initializeAttendanceTable()
		if err != nil {
			panic(err)
		}
	}
}

func GetDb() *sql.DB {
	return sqlDb
}
