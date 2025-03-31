package api

import (
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/db"
	rmq "github.com/ALTSKUF/ALTSKUF.Back.SquadData/rabbitmqclient"
  "github.com/gin-gonic/gin"

	"net/http"
	"log"
)

type Server struct {
  Db db.Db
	RabbitClient rmq.RabbitMQClient 
}

func Init(config *config.Config) (*Server, error) {
  db, err := db.Init(config)
  if err != nil {
    return nil, err
  }

  db.Migrate()

	rmq, err := rmq.Setup(config)
	if err != nil {
		return nil, err
	}

	return &Server{Db: db, RabbitClient: rmq}, nil
}

func (s Server) GetSquads(c *gin.Context) {
  squads, err := s.Db.GetAllSquads()
  if err != nil {
    c.Error(err)
    return
  }

  c.JSON(http.StatusOK, squads)
}

func (s Server) GetSquadById(c *gin.Context, squadId int) {
	response := s.Db.GetSquadById(squadId)
	if response.Error != nil {
		c.Error(response.Error)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (s Server) GetSquadMembersById(c *gin.Context, squadId int) {
	uuids, err := s.Db.GetSquadMembers(squadId)
	if err != nil {
		c.Error(err)
		return
	}

	log.Printf(" [*] Send request")
	log.Printf(" [*] Waiting for response")
	response := s.RabbitClient.GetUsersRPC(uuids)
	if response.Error != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, uuids)
}
