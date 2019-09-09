package jwt

import (
	"fmt"
	"net/http"
	"time"

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
		// 判断 token 是否有效
		if time.Now().Unix() > user.ValidBefore {
			token, err := util.GenerateToken(user.ID, user.Email)
			fmt.Println(token, err)
			if err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "StatusUnauthorized",
				})
			}
			c.Writer.Header().Set("Authorization", "Bearer "+token)
		}
		c.Next()
	}
}
