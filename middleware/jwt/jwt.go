package jwt

import (
	"fmt"
	"net/http"

	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		// Parse the header to get the token part
		var token string
		fmt.Sscanf(header, "Bearer %s", &token)
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "StatusBadRequest",
			})
			c.Abort()
			return
		}
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
