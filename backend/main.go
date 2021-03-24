package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/test", func(con *gin.Context) {
		con.JSON(http.StatusOK, gin.H{
			"response": "This api link is temporary.",
		})
	})

	groupSorts(router)

	router.Run(":3001")
}
