package test

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	v1 "github.com/handnotes/handnote-server/api/v1"
	"github.com/handnotes/handnote-server/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	InitialRouter()
	os.Exit(m.Run())
}

type AnyTime struct{}

func TestRegister(t *testing.T) {

	t.Run("it should be failed when no request body", func(t *testing.T) {
		db, _ := SetupTestDB()
		defer db.Close()

		code, body := HttpPost("/v1/auth/register", nil)
		assert.Equal(t, http.StatusBadRequest, code)
		var res v1.ResponseWithMessage
		_ = json.Unmarshal(body, &res)
		assert.Equal(t, "表单校验失败", res.Message)
	})

	t.Run("it should be success when correct params", func(t *testing.T) {
		db, mock := SetupTestDB()
		defer db.Close()

		mock.ExpectQuery(`^SELECT \* FROM "users"*`).
			WithArgs("mutoe@foxmail.com").
			WillReturnRows(sqlmock.NewRows([]string{}))
		mock.ExpectBegin()
		mock.ExpectExec(`^INSERT (.+)`).
			WithArgs("mutoe@foxmail.com", "mutoe", sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(`^SELECT (.+)`).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{}))
		mock.ExpectCommit()

		mock.ExpectBegin()
		mock.ExpectExec(`^INSERT (.+)`).
			WithArgs("memo", 1).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		form := url.Values{
			"email":     []string{"mutoe@foxmail.com"},
			"user_name": []string{"mutoe"},
			"password":  []string{"123456"},
			"code":      []string{"123456"},
		}
		HttpPost("/v1/auth/register", form)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("it should be failed when duplicate email", func(t *testing.T) {
		db, mock := SetupTestDB()
		defer db.Close()

		db.Exec(`INSERT INTO "users" ("email", "user_name", "password") VALUES (?, ?, ?)`, "mutoe@foxmail.com", "mutoe", "123456")

		mock.ExpectQuery(`^SELECT \* FROM "users"*`).
			WithArgs("mutoe@foxmail.com").
			WillReturnRows(sqlmock.NewRows([]string{}))

		form := url.Values{
			"email":     []string{"mutoe@foxmail.com"},
			"user_name": []string{"mutoe"},
			"password":  []string{"123456"},
			"code":      []string{"123456"},
		}
		code, body := HttpPost("/v1/auth/register", form)

		assert.Equal(t, http.StatusBadRequest, code)
		var res v1.ResponseWithMessage
		_ = json.Unmarshal(body, &res)
		assert.Equal(t, "创建用户失败", res.Message)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("it should be failed when incorrect verification code", func(t *testing.T) {
		form := url.Values{
			"email":     []string{"mutoe@foxmail.com"},
			"user_name": []string{"mutoe"},
			"password":  []string{"123456"},
			"code":      []string{"0"},
		}
		code, body := HttpPost("/v1/auth/register", form)

		assert.Equal(t, http.StatusBadRequest, code)
		var res v1.ResponseWithMessage
		_ = json.Unmarshal(body, &res)
		assert.Equal(t, "验证码失效", res.Message)

	})

}

func TestLogin(t *testing.T) {
	t.Run("it should be success when login with correct params", func(t *testing.T) {
		db, mock := SetupTestDB()
		defer db.Close()

		mock.ExpectQuery(`^SELECT*`).
			WithArgs("mutoe@foxmail.com").
			WillReturnRows(sqlmock.NewRows([]string{}))

		form := url.Values{}
		form.Add("email", "mutoe@foxmail.com")
		form.Add("password", "123456")
		code, body := HttpPost("/v1/auth/login", form)

		assert.Equal(t, http.StatusOK, code)
		var res v1.AuthResponse
		_ = json.Unmarshal(body, &res)
		assert.IsType(t, "string", res.Body.AccessToken)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})

	t.Run("it should be failed when login with incorrect params", func(t *testing.T) {
		db, mock, _ := sqlmock.New()
		gormDB, _ := gorm.Open("sqlite3", db)
		models.DB = gormDB
		defer db.Close()

		mock.ExpectQuery(`^SELECT*`).
			WillReturnRows(sqlmock.NewRows([]string{}))

		form := url.Values{}
		form.Add("email", "mutoe@foxmail.com")
		form.Add("password", "1234567")
		code, body := HttpPost("/v1/auth/login", form)

		assert.Equal(t, 400, code)
		var res v1.ResponseWithMessage
		_ = json.Unmarshal(body, &res)
		assert.Equal(t, "邮箱和密码匹配不上", res.Message)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
