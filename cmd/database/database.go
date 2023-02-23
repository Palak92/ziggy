package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	//Capture connection properties.
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "password",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "coins",
		AllowNativePasswords: true,
	}
	// Get a database handle.

	//dsn := "palakbansal:plk@tcp(localhost:3305)/coins"
	//dsn := "root:password@tcp(127.0.0.1:3306)/coins?tls=false&authPlugin=mysql_native_password"
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("err while opening mysql %v", err)
	}
	defer db.Close()

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatalf("err while pinging mysql %v", pingErr)
	}
	fmt.Println("Connected!")
}
