package main

import (
  "github.com/gin-gonic/gin"
  "log"
  "net/http"
  "os"
)

func main() {
  router := gin.Default()

  user := os.Getenv("USER")
  password := os.Getenv("PASSWORD")

  var authorized *gin.RouterGroup
  if (user != "" && password != "") {
    log.Printf("* PROTECTED BY USER AND PASSWORD *")

    authorized = router.Group("/", gin.BasicAuth(gin.Accounts{
      user: password,
    }))
  } else {
    log.Printf("* FREE ACCESS *")

    authorized = router.Group("/")
  }

  // / endpoint
  authorized.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "go away!")
  })

  // /pdf endpoint
  authorized.GET("/pdf", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"authorized": "yes"})
  })

  router.Run()
}
