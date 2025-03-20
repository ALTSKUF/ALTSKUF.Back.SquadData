package apperror

type AppError int

const (
  DbOpenError AppError = iota
  DbTransactionError
  RMQConnectionOpenError
  RMQChannelOpenError
  RMQQueueDeclareError
  RMQStartConsumingError

  InvalidURLParamError
  SquadNotFoundError
)

func (err AppError) Error() string {
  var str string
  switch err {
  case DbOpenError:
    str = "Database: open error"
  case DbTransactionError:
    str = "Database: transaction error"
  case RMQConnectionOpenError:
    str = "RabbitMQ: connection open error"
  case RMQChannelOpenError:
    str = "RabbitMQ: channel open error"
  case RMQQueueDeclareError:
    str = "RabbitMQ: failed to declare RPC queue"
  case RMQStartConsumingError:
    str = "RabbitMQ: failed to start consuming"
  case InvalidURLParamError:
    str = "Invalid URL"
  case SquadNotFoundError:
    str = "Squad not found"
  }

  return str
}
