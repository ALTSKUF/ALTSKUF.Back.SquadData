package main

import (
	e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/db"
	m "github.com/ALTSKUF/ALTSKUF.Back.SquadData/middleware"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/models"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/mqclient"
	"github.com/gin-gonic/gin"

	"log"
	"strconv"
)

func main() {
  config := config.Default()

  db, err := db.Init(config)
  if err != nil {
    log.Fatal(err)
  }

  db.AutoMigrate(&models.Squad{})
  db.AutoMigrate(&models.SquadMember{})

  rmq, err := mqclient.Setup(config)
  if err != nil {
    log.Fatal(err)
  }
  defer rmq.Close()

  gin.SetMode(config.AppProfile)
  router := gin.Default()
  router.Use(m.ErrorCatchMiddleware())

  router.GET("/squads/:id", func(c *gin.Context) {
    param := c.Param("id")
    squad_id, err := strconv.Atoi(param)
    if err != nil {
      c.Error(e.InvalidURLParamError)
        return
    }

    squad_info, err := db.GetSquadInfo(squad_id)
    if err != nil {
      c.Error(err)
      return
    }
    if squad_info == nil {
      response := gin.H{
        "info": "Squad doesn't exist",
      }
      c.JSON(404, response)
      return
    }

    uuids, err := db.GetSquadMembers(squad_id)
    if err != nil {
      c.Error(err)
      return
    }

    if len(uuids) == 0 {
      response := gin.H{
        "info": "Squad is empty",
      }

      c.JSON(200, response)
      return
    }

    response, err := rmq.GetUsersRPC(uuids)
    if err != nil {
      c.Error(err)
      return
    }

    c.JSON(200, response)
  })

  router.Run(config.AppAddress)
}
