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
	Email    string `json:"email" binding:"required,email"`
	UserName string `json:"user_name"`
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

// SignUpRequest 用户注册请求参数
// swagger:parameters signUpRequest
type SignUpRequest struct {
	// in: body
	Body SignUpForm
}

// SignResponse 用户注册/登录响应参数
// swagger:response signResponse
type SignResponse struct {
	// in: body
	Body struct {
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
}

// SignUpForm 用户注册表单
type SignUpForm struct {
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone" binding:"required"`
	UserName  string    `json:"user_name" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Address   string    `json:"address"`
	Gender    int8      `json:"gender" binding:"required"`
	Birth     time.Time `json:"birth"`
	AvatarURL string    `json:"avatar_url"`
	Code      int       `json:"code"`
}

// SignUp swagger:route POST /signUp signUpRequest
//
// 用户注册
//
//      Schemes: http, https
//
//      Responses:
//        200: signResponse
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
		Phone:     request.Phone,
		Email:     request.Email,
		UserName:  request.UserName,
		Password:  request.Password,
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

// SignInFormRequest 用户注册请求参数
// swagger:parameters signInRequest
type SignInFormRequest struct {
	// in: body
	Body SignInForm
}

// SignInForm 用户登录表单
type SignInForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignIn swagger:route POST /signIn signInRequest
//
// 用户登录
//
//      Schemes: http, https
//
//      Responses:
//        200: signResponse
func SignIn(c *gin.Context) {
	var request SignInForm
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

	c.JSON(http.StatusOK, gin.H{"token": token})
}
