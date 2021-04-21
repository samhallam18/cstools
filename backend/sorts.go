package main

import (
	"fmt"
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
	convertItems := []int{}
	for _, i := range items {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		convertItems = append(convertItems, j)
	}
	steps := [][]int{}
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
	con.Header("Access-Control-Allow-Origin", "*")
	items := strings.Split(con.PostForm("items"), ",")
	convertItems := []int{}
	for _, i := range items {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		convertItems = append(convertItems, j)
	}
	steps := [][]int{}
	for i := 1; i < len(convertItems); i++ {
		key := convertItems[i]
		j := i - 1
		for j >= 0 && convertItems[j] > key {
			convertItems[j+1] = convertItems[j]
			j = j - 1
		}
		convertItems[j+1] = key
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
	fmt.Println(steps)
	con.JSON(http.StatusOK, gin.H{
		"response": items,
		"steps":    steps,
	})
}
