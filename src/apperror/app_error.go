package apperror

type AppError int

const (
  DbOpenError AppError = iota
)

func (err AppError) Error() string {
  var str string
  switch err {
  case DbOpenError:
    str = "Database open error"
  }

  return str
}
