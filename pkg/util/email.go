package util

import (
	"bytes"
	"github.com/alecthomas/template"
	"github.com/gin-gonic/gin"
	"github.com/handnotes/handnote-server/pkg/setting"
	"log"
	"net/smtp"
	"strings"
)

func SendEmail(email string, code int) (err error) {
	emailTemplate := `To: {{.to}}
From: {{.from}}
Subject: 您的验证码是 {{.code}} [{{.appName}}]
Content-Type: text/html; charset=UTF-8

<p>尊敬的 {{.appName}} 用户您好，您的验证码是：<strong style="font-weight: bold;font-size: 160%;">{{.code}}</strong></p>

<p>验证码{{.validityPeriod}}分钟内有效，请不要将此验证码泄漏给他人，感谢您的使用！</p>
`
	smtpServer := setting.Email.SmtpServer
	var host, port string
	Unpack(strings.Split(smtpServer, ":"), &host, &port)

	auth := smtp.PlainAuth("", setting.Email.From, setting.Email.SmtpPassword, host)

	var smtpMessage bytes.Buffer
	var tmpl *template.Template
	tmpl, err = template.New("Email").Parse(emailTemplate)
	if err != nil {
		return
	}
	templateVariables := gin.H{
		"appName":        setting.App.Name,
		"to":             email,
		"from":           setting.Email.From,
		"code":           code,
		"validityPeriod": setting.Code.ValidityPeriod.Minutes(),
	}
	if err = tmpl.Execute(&smtpMessage, templateVariables); err != nil {
		return
	}

	if err = smtp.SendMail(smtpServer, auth, setting.Email.From, []string{email}, smtpMessage.Bytes()); err != nil {
		return
	}

	log.Printf("Email send success to '%s'", email)
	return nil
}
