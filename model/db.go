package model

import (
	"database/sql"
	"fmt"
)

var godoDB *sql.DB

func Init() {
	godoDB, _ = sql.Open("sqlite3", "./godo.db")
	if err == nil {
		fmt.Println("success!")
	}
}
