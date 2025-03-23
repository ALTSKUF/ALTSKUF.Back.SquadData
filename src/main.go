package main

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	m "github.com/ALTSKUF/ALTSKUF.Back.SquadData/middleware"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/app"
	"github.com/gin-gonic/gin"

	"log"
)

func main() {
  config := config.Default()

  app, err := app.Init(config)
  if err != nil {
    log.Fatal(err)
  }

  gin.SetMode(config.AppProfile)
  router := gin.Default()
  router.Use(m.ErrorCatchMiddleware())

  squads := router.Group("/squads")
  squads.GET("/", app.SquadsHandler)

  router.Run(config.AppAddress)
}
