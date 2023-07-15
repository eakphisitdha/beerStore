package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func Mariadb() *sql.DB {

	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/data")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Mariadb Connected")
	}
	return db
}
