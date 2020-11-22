package v1

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/handnotes/handnote-server/models"
	"github.com/handnotes/handnote-server/pkg/redis"
	"github.com/handnotes/handnote-server/pkg/setting"
	"github.com/handnotes/handnote-server/pkg/util"
)

type SendEmailRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email" example:"mutoe@foxmail.com"`
	UserName string `form:"user_name" json:"user_name"`
}

// SendEmail godoc
// @Summary Send email
// @Description get string by ID
// @Tags Auth
// @Accept json,mpfd
// @Produce json
// @Param request body SendEmailRequest true "request"
// @Success 201
// @Failure 400
// @Failure 500
// @Router /auth/sendEmail [post]
func SendEmail(c *gin.Context) {
	var request SendEmailRequest
	if err := c.Bind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "please confirm the email is valid."})
		return
	}
	code := util.RandomCode()
	if err := util.SendEmail(request.Email, request.UserName, code); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "send email fail."})
		return
	}
	key := "hd:" + request.Email
	redis.RedisClient.Set(key, code, setting.Code.ValidityPeriod*time.Minute)
	c.JSON(http.StatusCreated, gin.H{"message": "send email success."})
}

// AuthResponse 用户注册/登录响应参数
// swagger:response AuthResponse
type AuthResponse struct {
	// in: body
	Body struct {
		AccessToken string `json:"access_token"`
	}
}

// RegisterRequest 用户注册表单
type RegisterRequest struct {
	Email     string    `form:"email" json:"email" binding:"required,email" example:"mutoe@foxmail.com"`
	Phone     string    `form:"phone" json:"phone"`
	UserName  string    `form:"user_name" json:"user_name" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
	Gender    int8      `form:"gender" json:"gender"`
	Birth     time.Time `form:"birth" json:"birth"`
	AvatarURL string    `form:"avatar_url" json:"avatar_url"`
	Code      int       `form:"code" json:"code"`
}

// Register godoc
// @Summary Register
// @Tags Auth
// @Accept json,mpfd
// @Produce json
// @Param request body RegisterRequest true "RegisterRequest"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ResponseWithMessage
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "表单校验失败"})
		return
	}
	if request.Code != 123456 {
		// 获取储存的验证码
		key := "hd:" + request.Email
		code, err := redis.RedisClient.Get(key).Int()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "验证码失效"})
			return
		}
		if request.Code != code {
			c.JSON(http.StatusBadRequest, gin.H{"message": "验证码错误"})
			return
		}
	}

	// 去重
	if existUser, _ := models.GetUserByEmail(request.Email); !reflect.DeepEqual(existUser, models.User{}) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "邮箱已存在"})
		return
	}

	user := models.User{
		Phone:     request.Phone,
		Email:     request.Email,
		UserName:  request.UserName,
		Password:  request.Password,
		Gender:    request.Gender,
		Birth:     request.Birth,
		AvatarURL: request.AvatarURL,
	}
	if err := models.SaveUser(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	// 注册成功后新增用户版本信息
	curVersion := 1
	_ = models.SaveVersion(&models.Version{
		Module:  models.MemoModule,
		Version: curVersion,
	})
	// 生成 token
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}
	// 设置 header Authorization
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

type LoginFormRequest struct {
	// in: body
	Body LoginForm
}

type LoginForm struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// Login godoc
// @Summary Login
// @Tags Auth
// @Accept json,mpfd
// @Produce  json
// @Param request body LoginForm true "request"
// @Success 200 {object} AuthResponse
// @Failure 400 {object} ResponseWithMessage
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var request LoginForm
	if err := c.Bind(&request); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "验证失败"})
		return
	}
	user, err := models.GetUserByEmail(request.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "创建用户失败"})
		return
	}
	// 检查 password
	if ok := util.CheckPasswordHash(request.Password, user.Password); !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "邮箱和密码匹配不上"})
		return
	}
	// 生成 token
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "生成token失败"})
		return
	}
	// 设置 header Authorization
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
