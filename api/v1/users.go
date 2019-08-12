package v1

import (
	"net/http"

	"e.coding.net/handnote/handnote/models"
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
