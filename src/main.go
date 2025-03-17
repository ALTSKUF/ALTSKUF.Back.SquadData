package main

import (
	"log"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/db"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/models"
	"github.com/gin-gonic/gin"
)

func main() {
  config := config.Default()

  db, err := db.InitDb(config)
  if err != nil {
    log.Fatal(err)
  }

  db.AutoMigrate(&models.Squad{})
  db.AutoMigrate(&models.SquadMember{})

  gin.SetMode(config.AppProfile)
  router := gin.Default()

  router.GET("/", func(c *gin.Context) {
    data := gin.H{"hello": "world"}

    c.JSON(200, data)
  })

  router.Run(config.AppAddress)
}
