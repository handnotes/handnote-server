package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"e.coding.net/handnote/handnote/models"
	"github.com/gin-gonic/gin"
)

// UpdateUserForm 更新用户表单
type UpdateUserForm struct {
	UserName  string    `json:"user_name" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Address   string    `json:"address"`
	Gender    int8      `json:"gender" binding:"required"`
	Birth     time.Time `json:"birth"`
	AvatarURL string    `json:"avatar_url"`
}

// UpdateUser 创建用户
func UpdateUser(c *gin.Context) {
	var request UpdateUserForm
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ID错误"})
		return
	}
	user := models.User{
		ID:        uint(id),
		UserName:  request.UserName,
		Phone:     request.Phone,
		Password:  request.Password,
		Email:     request.Email,
		Address:   request.Address,
		Gender:    request.Gender,
		Birth:     request.Birth,
		AvatarURL: request.AvatarURL,
	}
	if err := models.SaveUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
