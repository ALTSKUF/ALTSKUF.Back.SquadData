package main

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  "fmt"
  "github.com/gin-gonic/gin"
)

func main() {
  fmt.Printf("Hello world")
  config := config.Default()

  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    data := gin.H{"hello": "world"}

    c.JSON(200, data)
  })

  router.Run(config.AppAddress)
}
