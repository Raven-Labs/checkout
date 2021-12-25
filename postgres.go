package main

import (
	"database/sql"
	"fmt"
    _ "github.com/lib/pq"
	"log"

)

func databaseinit(){
	// Set psql config from global conf
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Database.Host, config.Database.Port, config.Database.User, config.Database.Password, config.Database.DBName)
	
	log.Println("Connecting to DB")
	
	// Connect to db
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
        log.Fatal(err)
    }

	// Close DB
	defer db.Close()
	defer log.Println("Closing DB")

	err = db.Ping()
	if err != nil {
        log.Fatal(err)
    }	


	
	fmt.Println("ok")
}