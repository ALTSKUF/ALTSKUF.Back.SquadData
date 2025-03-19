package apperror

type AppError int

const (
  DbOpenError AppError = iota
  RMQConnectionOpenError
  RMQChannelOpenError
  RMQStartConsumingError
)

func (err AppError) Error() string {
  var str string
  switch err {
  case DbOpenError:
    str = "Database: open error"
  case RMQConnectionOpenError:
    str = "RabbitMQ: connection open error"
  case RMQChannelOpenError:
    str = "RabbitMQ: channel open error"
  case RMQStartConsumingError:
    str = "RabbitMQ: failed to start consuming"
  }

  return str
}
