package apperror

type AppError int

const (
  DbOpenError AppError = iota
  DbTransactionError
  DbSquadNotFoundError
  RMQConnectionOpenError
  RMQChannelOpenError
  RMQQueueDeclareError
  RMQStartConsumingError
  RMQInvalidResponse

  InvalidURLParamError
)

func (err AppError) Error() string {
  var str string
  switch err {
  case DbOpenError:
    str = "Database: open error"
  case DbTransactionError:
    str = "Database: transaction error"
	case DbSquadNotFoundError:
		str = "Database: squad not found"	
  case RMQConnectionOpenError:
    str = "RabbitMQ: connection open error"
  case RMQChannelOpenError:
    str = "RabbitMQ: channel open error"
  case RMQQueueDeclareError:
    str = "RabbitMQ: failed to declare RPC queue"
  case RMQStartConsumingError:
    str = "RabbitMQ: failed to start consuming"
  case RMQInvalidResponse:
    str = "RabbitMQ: invalid response"
  case InvalidURLParamError:
    str = "Invalid URL"
  }

  return str
}
