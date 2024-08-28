package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Init() {
	connect := "user=postgres password=12345 dbname=postgres sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		log.Fatal(err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение к БД")

}
