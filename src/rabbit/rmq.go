package rabbit

import (
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/config"
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  amqp "github.com/rabbitmq/amqp091-go"
  "fmt"
  "log"
)

type RMQ struct {
  Connection *amqp.Connection
  receiveChannel *amqp.Channel
  sendChannel *amqp.Channel
}

func Setup(config *config.Config) (*RMQ, error) {
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

  rch, err := conn.Channel()
  if err != nil {
    return nil, e.RMQChannelOpenError
  }

  sch, err := conn.Channel()
  if err != nil {
    return nil, e.RMQChannelOpenError
  }

  return &RMQ{conn, rch, sch}, nil
}

func (rmq *RMQ) StartConsuming(name string) error { // call as goroutine
  q, err := rmq.receiveChannel.QueueDeclare(
    name,
    false,
    false,
    true,
    false,
    nil,
  )

  msgs, err := rmq.consume(q.Name, "", true, true, false, false)
  if err != nil {
    return err
  }

  for msg := range msgs {
    log.Printf("Received message: %s", msg.Body)
  }

  return nil
}

func (rmq *RMQ) consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool) (<-chan amqp.Delivery, error) {
  delivery, err := rmq.receiveChannel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, nil) 
  if err != nil {
    return nil, e.RMQStartConsumingError
  }

  return delivery, nil
}

func (rmq *RMQ) Close() {
  rmq.receiveChannel.Close()
  rmq.sendChannel.Close()
  rmq.Connection.Close()
}
