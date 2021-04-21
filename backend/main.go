package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/test", func(con *gin.Context) {
		con.Header("Access-Control-Allow-Origin", "http://example.com")
		con.Header("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, OPTIONS")
		con.JSON(http.StatusOK, gin.H{
			"response": "This api link is temporary.",
		})
	})

	groupSorts(router)
	groupSearches(router)

	router.Use(cors.Default())
	router.Run(":3001")
}
