package util

import (
	"fmt"
	"github.com/handnotes/handnote-server/pkg/setting"
	"net/smtp"
	"os"
	"strings"
)

// SendEmail 发送验证码邮件方法
func SendEmail(email string, userName string, code int) error {
	var host, port string
	smtpServer := os.Getenv("SMTP_SERVER")
	Unpack(strings.Split(smtpServer, ":"), &host, &port)

	auth := smtp.PlainAuth("", setting.Email.From, os.Getenv("SMTP_PASSWORD"), host)
	contentType := "Content-Type: text/plain; charset=UTF-8"

	htmlContent := fmt.Sprintf("<p>尊敬的Handnote用户%s：</p><p>您好!</p><p>您的验证码是：%d。</p>", userName, code)
	msg := []byte("To: " + email + "\r\nFrom: " + setting.Email.From + "\r\nSubject: " + setting.Email.Subject + "\r\n" + contentType + "\r\n\r\n" + htmlContent)
	err := smtp.SendMail(smtpServer, auth, setting.Email.From, []string{email}, msg)
	return err
}
