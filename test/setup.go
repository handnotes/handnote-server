package test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/handnotes/handnote-server/models"
	"github.com/handnotes/handnote-server/routes"
	"github.com/jinzhu/gorm"
)

var (
	router *gin.Engine
)

func InitialRouter() {
	router = routes.SetupRouter()
}

func SetupTestDB() (gormDB *gorm.DB, mock sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, _ = gorm.Open("sqlite3", db)

	models.DB = gormDB
	return
}

func GenerateUserSeeder() {
	user := models.User{
		ID:       1,
		Email:    "root@admin.com",
		UserName: "root",
		Password: "123456",
	}
	_ = models.SaveUser(&user)
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
