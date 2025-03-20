package middleware

import (
  e "github.com/ALTSKUF/ALTSKUF.Back.SquadData/apperror"
  "github.com/gin-gonic/gin"
)

func ErrorCatchMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Next()

    if len(c.Errors) > 0 {
      err := c.Errors.Last().Err
      switch err{
      case e.InvalidURLParamError:
        c.JSON(400, gin.H{"error": "Invalid url"})
      default:
        c.JSON(500, gin.H{"error": "Internal server error"})
      }
    }
  }
}
