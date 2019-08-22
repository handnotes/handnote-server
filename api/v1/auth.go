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

// SendEmailRequest 发送邮件请求结构
type SendEmailRequest struct {
	UserName string `json:"user_name"`
	Email    string `json:"email" binding:"required,email"`
}

// SendEmail 发送邮件
func SendEmail(c *gin.Context) {
	var request SendEmailRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please confirm the email is valid."})
		return
	}
	fmt.Println(request)
	code := util.RandomCode()
	if err := util.SendEmail(request.Email, request.UserName, code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "send email fail."})
		return
	}
	key := "hd:" + request.Email
	redis.RedisClient.Set(key, code, setting.Code.ValidityPeriod*time.Minute)
	c.JSON(http.StatusOK, gin.H{"message": "send email success."})
}

// SignUpForm 用户注册表单
type SignUpForm struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"user_name" binding:"required"`
	Phone     string    `json:"phone" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Address   string    `json:"address"`
	Gender    int8      `json:"gender" binding:"required"`
	Birth     time.Time `json:"birth"`
	AvatarURL string    `json:"avatar_url"`
	Code      int       `json:"code"`
}

// SignUp 用户注册
func SignUp(c *gin.Context) {
	var request SignUpForm
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	// 获取储存的验证码
	key := "hd:" + request.Email
	code, err := redis.RedisClient.Get(key).Int()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证码失效"})
		return
	}
	if request.Code != code {
		c.JSON(http.StatusBadRequest, gin.H{"message": "请输入正确的验证码"})
		return
	}
	user := models.User{
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
	// 生成 token
	fmt.Println(user.ID)
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}