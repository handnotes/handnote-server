package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/handnotes/handnote-server/models"
	"github.com/handnotes/handnote-server/routes"
)

var (
	router     *gin.Engine
	ApiBaseUrl = "/api/v1"
)

func InitialRouter() {
	router = routes.SetupRouter()
}

func SetupTestDB() (gormDB *gorm.DB, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ = gorm.Open("sqlite3", db)

	gormDB.LogMode(true)
	models.DB = gormDB

	return
}

func HttpGet(uri string) (code int, body []byte) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", uri, nil)
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()

	code = result.StatusCode
	body, _ = ioutil.ReadAll(result.Body)
	return
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
