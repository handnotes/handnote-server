package test

import (
	v1 "e.coding.net/handnote/handnote/api/v1"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func init() {
	InitialRouter()
}

func TestRegister(t *testing.T) {

	t.Run("it should be failed when no request body", func(t *testing.T) {
		form := url.Values{}
		code, _ := HttpPost("/v1/auth/register", form)

		assert.Equal(t, 400, code)
	})

	t.Run("it should be success when correct params", func(t *testing.T) {
		form := url.Values{}
		form.Add("email", "mutoe@foxmail.com")
		form.Add("user_name", "mutoe1")
		form.Add("password", "123456")
		form.Add("code", "123456")
		code, body := HttpPost("/v1/auth/register", form)
		var res v1.AuthResponse
		_ = json.Unmarshal(body, res)

		assert.Equal(t, 200, code)
		assert.IsType(t, "string", res.Body.AccessToken)
	})

	t.Run("it should be failed when incorrect verification code", func(t *testing.T) {
		form := url.Values{}
		form.Add("email", "mutoe@foxmail.com")
		form.Add("user_name", "mutoe")
		form.Add("password", "123456")
		form.Add("code", "0")
		code, body := HttpPost("/v1/auth/register", form)

		assert.Equal(t, 400, code)
		var res v1.ResponseWithMessage
		_ = json.Unmarshal(body, &res)
		assert.Equal(t, "验证码失效", res.Message)
	})

}
