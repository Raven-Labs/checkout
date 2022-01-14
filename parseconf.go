package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	listenAddress string
	DefaultAdmin  struct {
		Password string
	}
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
	}
}

func getConfig() {
	log.Println("Grabbing config")
	content, err := ioutil.ReadFile("./config.json")

	//check for errors opening file
	if err != nil {
		log.Println("Failed to open config file. This usually means you need to rename config.example.json to config.json and customise the values.")
		log.Fatal("Error when opening config file: ", err)
	}

	// Unmarshal json
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Fatal("Error when Unmarshaling config json: ", err)
	}

	// Check for default credentials
	if config.DefaultAdmin.Password == "changeme" {
		log.Fatal("Default admin password used. Refusing to continue.")
	}

}
