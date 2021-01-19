package config

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	InfoColor ="\033[1;34m%s\033[0m"
)

func initializeDatabase() *sql.DB {
	fmt.Println(InfoColor, "Initialization of database")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/go")

    if err != nil {
        panic(err.Error())
	}
	
	return db
}