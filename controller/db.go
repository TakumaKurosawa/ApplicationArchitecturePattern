package controller

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/todo?parseTime=true")
	if err != nil {
		panic(err)
	}
}
