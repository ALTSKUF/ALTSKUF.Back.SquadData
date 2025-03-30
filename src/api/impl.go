package api

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/db"
  "github.com/gin-gonic/gin"

	"net/http"
)

type Server struct {
  Db db.Db
}

func Init(config *config.Config) (*Server, error) {
  db, err := db.Init(config)
  if err != nil {
    return nil, err
  }

  db.Migrate()

  return &Server{Db: db}, nil
}

func (s Server) GetSquads(c *gin.Context) {
  squads, err := s.Db.GetAllSquads()
  if err != nil {
    c.Error(err)
    return
  }

  c.JSON(http.StatusOK, squads)
}

func (s Server) GetSquadById(c *gin.Context, squad_id int) {
	response := s.Db.GetSquadById(squad_id)
	if response.Error != nil {
		c.Error(response.Error)
		return
	}

	c.JSON(http.StatusOK, response)
}
