package rabbitmqclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
	"github.com/ALTSKUF/ALTSKUF.Back.SquadData/dto"
	u "github.com/ALTSKUF/ALTSKUF.Back.SquadData/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
)

// Test microservice with users database, in real integration tests should be replaced
// Returns two users for non-empty list of uuids and empty response for empty list of uuids
func RabbitMQServer() {
  config := config.Default()

  dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.RMQUser, config.RMQPassword, config.RMQHost, config.RMQPort)
  log.Printf("Server connecting to RabbitMQ")
  conn, err := amqp.Dial(dsn)
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()

  ch, err := conn.Channel()
  if err != nil {
    log.Fatal(err)
  }
  defer ch.Close()

  q, err := ch.QueueDeclare(
    "rpc_queue",
    false,
    false,
    false,
    false,
    nil,
  )
  if err != nil {
    log.Fatal(err)
  }

  delivery, err := ch.Consume(
    q.Name,
    "",
    true,
    false,
    false,
    false,
    nil,
  )

  log.Printf(" [+] Start processing")
  for rpc := range delivery {
    var sendUUIDS dto.SendUUIDS      
    var response dto.GetUsersResponse

    err := json.Unmarshal(rpc.Body, &sendUUIDS)
    if err != nil {
      response = dto.GetUsersResponse{
        Error: errors.New("Invalid argument"),
      }
    } else {
      log.Printf(" [+] Got UUIDs: %s", sendUUIDS)

      if len(sendUUIDS.UUIDS) == 0 {
        response = dto.GetUsersResponse {}
      } else {
        sendUsers := []dto.User {
          {FullName: "John Doe", Group: "secret", Role: "manager"},
          {FullName: "Doe John", Group: "non secret", Role: "director"},
        }

        response = dto.GetUsersResponse {
          Users: &sendUsers,
        }
      }
    }

    body, err := json.Marshal(response)
    if err != nil {
      log.Fatal(err)
    }

    err = ch.Publish(
      "",
      rpc.ReplyTo,
      false,
      false,
      amqp.Publishing{
        ContentType: "application/json",
        CorrelationId: rpc.CorrelationId,
        Body: body,
      },
    )

    if err != nil {
      log.Fatal(err)
    }
    log.Printf(" [+] Sent %s", response.Users)
  }
}

// Test GetUsersRPC with empty list of uuids
// Response must be empty
func TestGetUsersRPCEmpty(t *testing.T) {
  u.LongTest(t)

  config := config.Default()
  client, err := Setup(config)
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  // remove in future
  go RabbitMQServer()

  uuids := make([]uuid.UUID, 0)

  log.Printf(" [*] Waiting for response")
  response := client.GetUsersRPC(uuids)
  log.Printf(" [*] Got users: %s", response.Users)

  assert.Condition(t, func () bool { 
    return response.Users == nil && response.Error == nil
  })
}

// Test GetUsersRPC with non-empty list of uuids
// Response must contain list of two users
func TestGetUsersRPCNotEmpty(t *testing.T) {
  u.LongTest(t)

  config := config.Default()

  client, err := Setup(config)
  if err != nil {
    log.Fatal(err)
  }
  defer client.Close()

  uuids := []uuid.UUID {
    uuid.New(),
    uuid.New(),
  }

  // remove in future
  go RabbitMQServer()

  log.Printf("[*] Waiting for response")

  response := client.GetUsersRPC(uuids)

  log.Printf(" [*] Got users: %s", response.Users)

  if response.Error != nil {
    log.Printf(" [*] Error: %s", response.Error)
  }

  assert.Condition(t, func () bool { 
    return response.Users != nil && len(*response.Users) == 2 && response.Error == nil
  })
}
