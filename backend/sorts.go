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
	items := strings.Split(con.PostForm("items"), ",")
	steps := [][]string{}
	for i := 0; i < len(items)-1; i++ {
		swapped := false
		for j := 1; j < len(items)-i; j++ {
			if items[j-1] > items[j] {
				temp := items[j-1]
				items[j-1] = items[j]
				items[j] = temp
				swapped = true
			}
			tempCopy := make([]string, len(items))
			copy(tempCopy, items)
			steps = append(steps, tempCopy)
		}
		if !swapped {
			break
		}
	}
	con.JSON(http.StatusOK, gin.H{
		"response": items,
		"steps":    steps,
	})

}
