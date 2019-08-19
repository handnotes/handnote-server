package util

import (
	"math/rand"
	"time"

	"e.coding.net/handnote/handnote/pkg/setting"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomString A helper function to generate random string with length n.
func RandomString(n int) string {
	str := make([]rune, n)
	for i := range str {
		str[i] = letters[rand.Intn(len(letters))]
	}
	return string(str)
}

// RandomCode A helper function to generate random verification code.
func RandomCode() int {
	min := setting.Code.Min
	max := setting.Code.Max + 1
	return rand.Intn(max-min) + min
}

// GeneratePassword returns the bcrypt hash of the password with default cost.
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash check the given hashed password right or not. Returns true
// on success, or false on failure.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

const signKey = "this is a sign key to generate token..."

// GenerateToken returns a signed token by jwt.
func GenerateToken(id uint) string {
	jwtToken := jwt.New(jwt.GetSigningMethod("HS256"))
	// Set some claims
	jwtToken.Claims = jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	// Sign and get the complete encoded token as a string
	token, _ := jwtToken.SignedString([]byte(signKey))
	return token
}
