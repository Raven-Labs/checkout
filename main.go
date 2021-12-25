package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func test(c *gin.Context){
	dbname := getConfig("dbname")
	c.IndentedJSON(http.StatusOK, dbname)
}

func main(){
	router := gin.Default()
    router.GET("/test", test)

    router.Run("localhost:8080")
}