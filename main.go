package main

import (
	"github/adekang/gin-demo/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	// db := common.InitDB()
	// defer db.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r = CollectRoute(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
