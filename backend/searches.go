package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func groupSearches(route *gin.Engine) {
	searches := route.Group("/searches")

	searches.POST("/binary", binarySearch)
}

func binarySearch(con *gin.Context) {
	con.Header("Access-Control-Allow-Origin", "*")
	searchValue, _ := strconv.Atoi(con.PostForm("searchValue"))
	items := []int{}

	for _, i := range strings.Split(con.PostForm("items"), ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		items = append(items, j)
	}

	steps := [][]int{}
	start := 0
	end := len(items) - 1
	var mid int
	found := false

	for start <= end && !found {
		mid = (start + end) / 2
		steps = append(steps, items[start:end+1])
		fmt.Println(steps)

		if items[mid] == searchValue {
			found = true
		} else if items[mid] < searchValue {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	if found {
		con.JSON(http.StatusOK, gin.H{
			"response": "Item found at" + strconv.Itoa(mid),
			"steps":    steps,
		})
	} else {
		con.JSON(http.StatusOK, gin.H{
			"response": "The item you searched for was not found.",
		})
	}
}
