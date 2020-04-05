package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Router() *gin.Engine {
	router := gin.Default()

	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")

	var authorized *gin.RouterGroup
	if user != "" && password != "" {
		fmt.Println("* PROTECTED BY USER AND PASSWORD *")

		authorized = router.Group("/", gin.BasicAuth(gin.Accounts{
			user: password,
		}))
	} else {
		fmt.Println("* OPEN ACCESS *")

		authorized = router.Group("/")
	}

	// / open endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "你好\nHello\nHola")
	})

	// /pdf protected endpoint
	authorized.GET("/pdf", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"authorized": "yes"})
	})

	return router
}
