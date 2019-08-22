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

// SendEmail 发送邮件
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

// CreateUser 更新用户信息
func CreateUser(c *gin.Context) {
	var user models.UserForm
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	// // 获取储存的验证码
	// key := "hd:" + user.Email
	// code, err := redis.RedisClient.Get(key).Int()
	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "验证码失效"})
	// 	return
	// }
	// if user.Code != code {
	// 	c.JSON(http.StatusBadRequest, gin.H{"message": "请输入正确的验证码"})
	// 	return
	// }
	if err := models.SaveUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	// 生成 token
	fmt.Println(user.ID)
	token, err := util.GenerateToken(user.ID, user.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
