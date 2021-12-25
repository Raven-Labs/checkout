package main
 
import (
    "encoding/json"
    "io/ioutil"
    "log"
)
 
type Config struct {
    Database struct{
		Host string
		Port string
		User string
		Password string
		DBName string
	}
}
 
func getConfig(entry string) string {
    content, err := ioutil.ReadFile("./config.json")

	//check for errors opening file
	if err != nil {
        log.Fatal("Error when opening config file: ", err)
    }
 
	// Unmarshal json
    var config Config
    err = json.Unmarshal(content, &config)
    if err != nil {
        log.Fatal("Error when Unmarshaling config json: ", err)
    }
 
	switch entry{
		case "dbport":
			return config.Database.Port
		case "dbname":
			return config.Database.DBName
		case "dbuser":
			return config.Database.User
		case "dbhost":
			return config.Database.Host
		case "dbpassword":
			return config.Database.Password
	}
	return "err"
}
