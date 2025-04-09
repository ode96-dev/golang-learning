package main

import "github.com/gin-gonic/gin"

/*
Gin
- lightweight, fast, and the most popular web frameworks in Go
features:
- fast
- middleware support
- routing with parameters and groups
- JSON binding and validation
*/

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "ping"})
	})

	r.Run() // defaults to 8080
}
