package app

import "github.com/gin-gonic/gin"

func (app *App) SquadsHandler(c *gin.Context) {
  squads, err := app.Db.GetAllSquads()
  if err != nil {
    c.Error(err)
    return
  }

  c.JSON(200, gin.H{
    "squads": squads,
  })
}
