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
			c.Error(e.InvalidURLParamError) // For now, the only parsing error that can happen is error when squad_id is invalid
		},
	})

  s := http.Server{
    Handler: router,
    Addr: config.AppAddress,
  }

  log.Fatal(s.ListenAndServe())
}
