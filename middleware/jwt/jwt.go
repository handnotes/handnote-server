package jwt

import (
	"net/http"

	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200
		token := c.Query("token")
		if token == "" {
			code = 400
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = 20002
				default:
					code = 20001
				}
			}
		}
		if code != 200 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  "test",
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
