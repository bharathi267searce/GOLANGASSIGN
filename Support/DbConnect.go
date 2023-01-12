package Support

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func DbConnect() {
	var err error
	var db *sql.DB
	connStr := "postgres://postgres:anbu@localhost/postgres?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		fmt.Println("hlo")
		panic(err)
	}
	DB = db
	fmt.Printf("\nSuccessfully connected to database!\n")
	// defer db.Close()
	// defer DB.Close()

}
