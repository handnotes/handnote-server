package v1

import (
	"fmt"
	"net/http"
	"time"

	"e.coding.net/handnote/handnote/models"
	"e.coding.net/handnote/handnote/pkg/redis"
	"e.coding.net/handnote/handnote/pkg/setting"
	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// SendEmail send email from handnote to destination email address.
func SendEmail(c *gin.Context) {
	var user models.UserEmail
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please confirm the email is valid."})
		return
	}
	fmt.Println(user)
	code := util.RandomCode()
	if err := util.SendEmail(user.Email, user.UserName, code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "send email fail."})
		return
	}
	key := "hd:" + user.Email
	redis.RedisClient.Set(key, code, setting.Code.ValidityPeriod*time.Minute)
	c.JSON(http.StatusOK, gin.H{"message": "send email success."})
}
