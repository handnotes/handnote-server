package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello is a test demo.
func Hello(c *gin.Context) {
	c.String(http.StatusOK, "Hello world.")
}
