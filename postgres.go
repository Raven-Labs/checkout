package main

import (
	"fmt"
)

func databaseinit(){
	dbname := getConfig("dbname")
	fmt.Println(dbname)
}