package v1

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Hello is a test demo.
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello world.")
}

// SendEmail send email from handnote to destination email address.
func SendEmail(c *gin.Context) {
	from := mail.NewEmail("Handnote", "no_reply@gethin.cn")
	subject := "Handnote注册验证码"
	to := mail.NewEmail("gethin.yan", "gethin.yan@foxmail.com")
	plainTextContent := "Handnote注册验证码"
	htmlContent := "<p>尊敬的Handnote用户：</p><p>您好!</p><p>您的注册验证码是：123456。</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	fmt.Println(os.Getenv("SENDGRID_API_KEY"))
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
