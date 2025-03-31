package rabbitmqclient

import (
	e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/schemas"
	u "github.com/ALTSKUF/ALTSKUF.Back.SquadData/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"fmt"
  "log"
)

type RabbitMQClient interface {
  GetUsersRPC([]uuid.UUID) schemas.GetUsersResponse
}

type Client struct {
  connection *amqp.Connection
  channel *amqp.Channel
}

func Setup(config *config.Config) (*Client, error) {
  url := fmt.Sprintf("amqp://%s:%s@%s:%s/", 
    config.RMQUser, 
    config.RMQPassword, 
    config.RMQHost, 
    config.RMQPort, 
  ) 

  log.Printf("Connecting to RabbitMQ on %s", url)
  conn, err := amqp.Dial(url)
  if err != nil {
    log.Printf("Error: %s", err)
    return nil, e.RMQConnectionOpenError
  }

  ch, err := conn.Channel()
  if err != nil {
    return nil, e.RMQChannelOpenError
  }

  return &Client{conn, ch}, nil
}

func (rmq *Client) GetUsersRPC(uuids []uuid.UUID) schemas.GetUsersResponse {
  if len(uuids) == 0 {
    return schemas.GetUsersResponse{}
  } 

  q, err := rmq.channel.QueueDeclare(
    "",
    false,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    return schemas.GetUsersResponse{
      Error: e.RMQQueueDeclareError,
    }
  }

  corrId := u.RandomString(32)

  sendUUIDS := schemas.SendUUIDS{UUIDS: uuids}

  body, _ := json.Marshal(sendUUIDS)
  rmq.channel.Publish(
    "",
    "rpc_queue",
    true, 
    false, 
    amqp.Publishing{
      ContentType: "application/json", 
      CorrelationId: corrId,
      ReplyTo: q.Name,
      Body: body,
    }, 
  ) 

  msgs, err := rmq.channel.Consume(q.Name, "", true, true, false, false, nil)
  if err != nil {
    return schemas.GetUsersResponse{
      Error: e.RMQStartConsumingError,
    } 
  }

  var response schemas.GetUsersResponse
  log.Println(" [*] Get user RPC")
  for msg := range msgs {
    if msg.CorrelationId != corrId {
      continue
    } 

    err := json.Unmarshal(msg.Body, &response)
    if err != nil {
      return schemas.GetUsersResponse{
        Error: e.RMQInvalidResponse,
      } 
    }
    break
  }
  log.Println(" [*] End of user RPC")
  return response
}

func (rmq *Client) Close() {
  rmq.channel.Close()
  rmq.connection.Close()
}
