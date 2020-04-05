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
    log.Printf("* OPEN ACCESS *")

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

  router.Run()
}
