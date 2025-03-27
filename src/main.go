package main

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	m "github.com/ALTSKUF/ALTSKUF.Back.SquadData/middleware"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/api"
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

  api.RegisterHandlers(router, server)

  s := http.Server{
    Handler: router,
    Addr: config.AppAddress,
  }

  log.Fatal(s.ListenAndServe())
}
