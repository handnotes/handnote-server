package jwt

import (
	"fmt"
	"net/http"
	"strings"

	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		splitToken := strings.Split(token, "Bearer")
		if len(splitToken) != 2 || splitToken[1] == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "StatusBadRequest",
			})
			c.Abort()
			return
		}
		token = strings.TrimSpace(splitToken[1])
		fmt.Println(token)
		user, err := util.ParseToken(token)
		fmt.Println(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "StatusUnauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
