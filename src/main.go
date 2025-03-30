package main

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/api"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	m "github.com/ALTSKUF/ALTSKUF.Back.SquadData/middleware"
	e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
	"github.com/gin-gonic/gin"

	"log"
	"net/http"
)

func main() {
  config := config.Default()

  server, err := api.Init(config)
  if err != nil {
    log.Fatal(err)
  }

  gin.SetMode(config.AppProfile)
  router := gin.Default()
  router.Use(m.ErrorCatchMiddleware())

  api.RegisterHandlersWithOptions(router, server, api.GinServerOptions{
		ErrorHandler: func(c *gin.Context, err error, statusCode int) {
			c.Error(e.InvalidURLParamError) // For now, only error that can happen is error when parsing squad_id parameter in /squads/{squad_id} route
		},
	})

  s := http.Server{
    Handler: router,
    Addr: config.AppAddress,
  }

  log.Fatal(s.ListenAndServe())
}
