package middleware

import (
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "github.com/ALTSKUF/ALTSKUF.Back.SquadData/schemas"
  "github.com/gin-gonic/gin"

  "net/http"
)

func ErrorCatchMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()

    if len(c.Errors) > 0 {
      err := c.Errors.Last().Err
      var error schemas.Error

      switch err{
      case e.InvalidURLParamError:
        error.Error = "Invalid url parameter"
        c.JSON(http.StatusBadRequest, error)
      default:
        error.Error = "Internal server error"
        c.JSON(http.StatusInternalServerError, error)
      }
    }
  }
}
