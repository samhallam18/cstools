package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func groupSorts(route *gin.Engine) {
	sorts := route.Group("/sorts")

	sorts.POST("/bubble", bubbleSort)
}

func bubbleSort(con *gin.Context) {
	con.Header("Content-Type", "application/json")
	items := strings.Split(con.PostForm("items"), ",")
	for i := 1; i < len(items); i++ {
		key := items[i]
		j := i - 1
		for j >= 0 && items[j] > key {
			items[j+1] = items[j]
			j--
		}
		items[j+1] = key
	}
	con.JSON(http.StatusOK, gin.H{
		"response": items,
	})
}
