package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func groupSorts(route *gin.Engine) {
	sorts := route.Group("/sorts")

	sorts.POST("/bubble", bubbleSort)
	sorts.POST("/insertion", insertionSort)
}

func bubbleSort(con *gin.Context) {
	con.Header("Access-Control-Allow-Origin", "*")
	con.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
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

func insertionSort(con *gin.Context) {
	items := strings.Split(con.PostForm("items"), ",")
	steps := [][]string{}
	for i := 1; i < len(items); i++ {
		key := items[i]
		j := i - 1
		for j >= 0 && items[j] > key {
			items[j+1] = items[j]
			j = j - 1
		}
		items[j+1] = key
		tempCopy := make([]string, len(items))
		copy(tempCopy, items)
		steps = append(steps, tempCopy)
	}
	con.JSON(http.StatusOK, gin.H{
		"response": items,
		"steps":    steps,
	})
}
