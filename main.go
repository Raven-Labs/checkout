package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
)

var config Config


func test(c *gin.Context){
	c.IndentedJSON(http.StatusOK, config.Database.Port)
}

func main(){
	log.Println("Starting Checkout...")
	getConfig()
	databaseinit()
	router := gin.Default()
    router.GET("/test", test)

    router.Run("localhost:8080")
}