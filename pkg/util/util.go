package util

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/handnotes/handnote-server/pkg/setting"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString 生成随机字符串
func RandomString(n int) string {
	str := make([]rune, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

// RandomCode 生成随机验证码
func RandomCode() int {
	min := setting.Code.Min
	max := setting.Code.Max + 1
	return rand.Intn(max-min) + min
}

func GeneratePassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Unpack(s []string, vars ...*string) {
	for i, str := range s {
		*vars[i] = str
	}
}
