package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var config Config
var db *sql.DB

func api_test(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, config.listenAddress)
}

func api_auth(c *gin.Context) {
	type account struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var authAttempt account

	if err := c.BindJSON(&authAttempt); err != nil {
		return
	}
	result := authAdmin(authAttempt.Username, authAttempt.Password)
	if result == 200 {
		c.IndentedJSON(http.StatusOK, "Good")
	}
	c.IndentedJSON(http.StatusUnauthorized, "Unauthorised")

}

func main() {
	log.Println("Starting Checkout...")
	getConfig()
	databaseinit()
	initAdmin()
	router := gin.Default()
	router.GET("/test", api_test)
	router.POST("/auth", api_auth)
	listenAddress := config.listenAddress + ":8080"

	router.Run(listenAddress)
}
