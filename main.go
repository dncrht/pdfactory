package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
    "user": "password",
	}))

	// /pdf endpoint
	authorized.GET("/secrets", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"authorized": "yes"})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run()
}
