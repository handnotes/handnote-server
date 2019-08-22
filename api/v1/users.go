package v1

import (
	"fmt"
	"net/http"

	"e.coding.net/handnote/handnote/models"
	"github.com/gin-gonic/gin"
)

// GetUsers 获取用户列表方法
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

// UpdateUser 创建用户
func UpdateUser(c *gin.Context) {
	var user models.UserForm
	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	if err := models.SaveUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
