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
	searches.POST("/linear", linearSearch)
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

	steps := []string{}
	start := 0
	end := len(items) - 1
	var mid int
	found := false

	for start <= end && !found {
		mid = (start + end) / 2
		fmt.Println(steps)

		if items[mid] == searchValue {
			steps = append(steps, "Item found at: "+strconv.Itoa(mid))
			found = true
		} else if items[mid] < searchValue {
			steps = append(steps, "Not found at middle value: "+strconv.Itoa(mid))
			start = mid + 1
		} else {
			steps = append(steps, "Not found at middle value: "+strconv.Itoa(mid))
			end = mid - 1
		}
	}

	if found {
		con.JSON(http.StatusOK, gin.H{
			"response": "Item found at index " + strconv.Itoa(mid),
			"steps":    steps,
		})
	} else {
		con.JSON(http.StatusOK, gin.H{
			"response": "The item you searched for was not found.",
		})
	}
}

func linearSearch(con *gin.Context) {
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

	steps := []string{}
	for _, i := range items {
		if items[i] == searchValue {
			steps = append(steps, "Item found at: "+strconv.Itoa(i))
			con.JSON(http.StatusOK, gin.H{
				"response": "Item found at: " + strconv.Itoa(i),
				"steps":    steps,
			})
			return
		} else {
			steps = append(steps, "Item not found at: "+strconv.Itoa(i))
		}
	}
}
