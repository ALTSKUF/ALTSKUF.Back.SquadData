package main

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  "github.com/gin-gonic/gin"
)

func main() {
  config := config.Default()

  gin.SetMode(config.AppProfile)
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    data := gin.H{"hello": "world"}

    c.JSON(200, data)
  })

  router.Run(config.AppAddress)
}
