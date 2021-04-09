package main

import (
	"net/http"
	"reflect"
	"strconv"
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
	items := strings.Split(con.PostForm("items"), ",")
	datatype := con.PostForm("type")
	var convertItems nil
	var steps nil
	if datatype == "string" {
		convertItems = items
	} else if datatype == "integer" {
		convertItems == []int{}
		for _, i := range items {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			convertItems = append(convertItems, j)
		}
	}
	for i := 0; i < len(convertItems)-1; i++ {
		swapped := false
		for j := 1; j < len(convertItems)-i; j++ {
			if convertItems[j-1] > convertItems[j] {
				temp := convertItems[j-1]
				convertItems[j-1] = convertItems[j]
				convertItems[j] = temp
				swapped = true
			}
			tempCopy := make([]int, len(convertItems))
			copy(tempCopy, convertItems)
			if len(steps) > 0 {
				if !reflect.DeepEqual(tempCopy, steps[(len(steps)-1)]) {
					steps = append(steps, tempCopy)
				}
			} else {
				steps = append(steps, tempCopy)
			}
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
