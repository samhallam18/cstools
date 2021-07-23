package main

import "github.com/gin-gonic/gin"

func stacks(con *gin.Context) {
	con.Header("Access-Control-Allow-Origin", "*")
	currentQ := con.PostForm("currentQ")
}
