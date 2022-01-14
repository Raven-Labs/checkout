package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func databaseinit() {
	var err error

	// Set psql config from global conf
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)

	log.Println("Connecting to DB")

	// Connect to db
	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	// Close DB
	//defer db.Close()
	//defer log.Println("Closing DB")

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create tables

	tables := []string{"admins"}
	for _, table := range tables {
		fmt.Printf(table)
		_, err = db.Exec(fmt.Sprintf("CREATE TABLE %s;", table))
		if err != nil {
			log.Fatal("Failed to create table ", table, ". ", err)
		}
	}

	fmt.Println("ok")
}
