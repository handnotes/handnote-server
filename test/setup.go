package test

import (
	"e.coding.net/handnote/handnote/routes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

var (
	router *gin.Engine
)

func InitialRouter() {
	router = routes.SetupRouter()
}

func HttpGet(uri string) []byte {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()

	body, _ := ioutil.ReadAll(result.Body)
	return body
}

func HttpPost(uri string, form url.Values) (code int, body []byte) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, uri, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	result := w.Result()
	defer result.Body.Close()

	code = result.StatusCode
	body, _ = ioutil.ReadAll(result.Body)
	return
}
