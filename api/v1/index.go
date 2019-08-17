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
	from := mail.NewEmail("Example User Gmail", "gethin.yan@gmail.com")
	subject := "Sending with Twilio SendGrid is Fun"
	to := mail.NewEmail("Example User Foxmail", "gethin.yan@foxmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
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
