package v1

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"e.coding.net/handnote/handnote/models"
	"e.coding.net/handnote/handnote/pkg/redis"
	"e.coding.net/handnote/handnote/pkg/setting"
	"e.coding.net/handnote/handnote/pkg/util"
	"github.com/gin-gonic/gin"
)

// SendEmailRequest 发送邮件请求结构
// swagger:parameters sendEmailRequest
type SendEmailRequest struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	UserName string `form:"user_name" json:"user_name"`
}

// SendEmail swagger:route GET /auth/sendEmail sendEmailRequest
//
// 发送邮件
//
// 		Schemes: http, https
//
// 		Responses:
//      	200: AuthResponse

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

// RegisterRequest 用户注册请求参数
// swagger:parameters registerRequest
type RegisterRequest struct {
	// in: body
	Body RegisterForm
}

// AuthResponse 用户注册/登录响应参数
// swagger:response AuthResponse
type AuthResponse struct {
	// in: body
	Body struct {
		AccessToken string `json:"access_token"`
	}
}

// RegisterForm 用户注册表单
type RegisterForm struct {
	Email     string    `form:"email" json:"email" binding:"required,email"`
	Phone     string    `form:"phone" json:"phone"`
	UserName  string    `form:"user_name" json:"user_name" binding:"required"`
	Password  string    `form:"password" json:"password" binding:"required"`
	Gender    int8      `form:"gender" json:"gender"`
	Birth     time.Time `form:"birth" json:"birth"`
	AvatarURL string    `form:"avatar_url" json:"avatar_url"`
	Code      int       `form:"code" json:"code"`
}

// Register swagger:route POST /auth/register registerRequest
//
// 用户注册
//
//     Schemes: http, https
//
//     Responses:
//       200: AuthResponse
//    	 400: ResponseWithMessage
func Register(c *gin.Context) {
	var request RegisterForm
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
	fmt.Println(user.ID)
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

// LoginFormRequest 用户注册请求参数
// swagger:parameters loginRequest
type LoginFormRequest struct {
	// in: body
	Body LoginForm
}

// LoginForm 用户登录表单
type LoginForm struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// Login swagger:route POST /auth/login loginRequest
//
// 用户登录
//
//      Schemes: http, https
//
//      Responses:
//        200: AuthResponse
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
	fmt.Println(user.ID)
	token, err := util.GenerateToken(user.ID, user.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "生成token失败"})
		return
	}
	// 设置 header Authorization
	c.Writer.Header().Set("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
