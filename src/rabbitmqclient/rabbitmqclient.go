package rabbitmqclient

import (
	e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/dto"
	u "github.com/ALTSKUF/ALTSKUF.Back.SquadData/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"encoding/json"
	"fmt"
  "log"
)

type RMQClient struct {
  connection *amqp.Connection
  channel *amqp.Channel
}

func Setup(config *config.Config) (*RMQClient, error) {
  url := fmt.Sprintf("amqp://%s:%s@%s:%s/", 
    config.RMQUser, 
    config.RMQPassword, 
    config.RMQHost, 
    config.RMQPort, 
  ) 

  conn, err := amqp.Dial(url)
  if err != nil {
    return nil, e.RMQConnectionOpenError
  }

  ch, err := conn.Channel()
  if err != nil {
    return nil, e.RMQChannelOpenError
  }

  return &RMQClient{conn, ch}, nil
}

func (rmq *RMQClient) consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
  delivery, err := rmq.channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, nil) 
  if err != nil {
    return nil, e.RMQStartConsumingError
  }

  return delivery, nil
}

func (rmq *RMQClient) GetUsersRPC(uuids []uuid.UUID) (*dto.GetUserResponse, error) {
  q, err := rmq.channel.QueueDeclare(
    "",
    false,
    false,
    true,
    false,
    nil,
  )
  if err != nil {
    return nil, e.RMQQueueDeclareError
  }

  corrId := u.RandomString(32)

  body, _ := json.Marshal(uuids)
  rmq.channel.Publish(
    "",
    q.Name,
    true, 
    false, 
    amqp.Publishing{
      ContentType: "application/json", 
      CorrelationId: corrId,
      ReplyTo: q.Name,
      Body: body,
    }, 
  ) 

  msgs, err := rmq.consume(q.Name, "", true, true, false, false)
  if err != nil {
    return nil, e.RMQStartConsumingError
  }

  var response dto.GetUserResponse
  log.Println(" [*] Get user RPC")
  for msg := range msgs {
    if msg.CorrelationId != corrId {
      continue
    } 

    json.Unmarshal(msg.Body, &response)
    break
  }
  log.Println(" [*] End of user RPC")
  return &response, nil
}

func (rmq *RMQClient) Close() {
  rmq.channel.Close()
  rmq.connection.Close()
}
