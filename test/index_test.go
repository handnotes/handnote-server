package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"e.coding.net/handnote/handnote/routes"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	router := routes.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/hello", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "Hello world.", w.Body.String())
}
