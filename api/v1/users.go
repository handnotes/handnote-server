package v1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handnotes/handnote-server/models"
)

// UpdateUserForm 更新用户表单
type UpdateUserForm struct {
	Email     string    `form:"email" json:"email" binding:"required,email"`
	UserName  string    `form:"user_name" json:"user_name" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
	Phone     string    `form:"phone" json:"phone"`
	Gender    int8      `form:"gender" json:"gender"`
	Birth     time.Time `form:"birth" json:"birth"`
	AvatarURL string    `form:"avatar_url" json:"avatar_url"`
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
