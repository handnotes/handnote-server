package v1

import (
	"fmt"
	"net/http"

	"e.coding.net/handnote/handnote/models"
	"e.coding.net/handnote/handnote/pkg/redis"
	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表方法.
func GetUsers(c *gin.Context) {
	userList, err := models.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "抱歉未找到用户",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": userList,
	})
}

// CreateUser is aimed to add user.
func CreateUser(c *gin.Context) {
	var user models.UserForm
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	// 获取储存的验证码
	key := "hd:" + user.Email
	code, err := redis.RedisClient.Get(key).Int()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证码失效"})
		return
	}
	if user.Code != code {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请输入正确的验证码"})
		return
	}
	if user.Password, err = util.GeneratePassword(user.Password); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "密码加密失败"})
		return
	}
	if err := models.CreateUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	// 生成 token
	fmt.Println(user.ID)
	token := util.GenerateToken(user.ID)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
