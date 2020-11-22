package setting

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

// AppSection app.yaml server 配置
type AppSection struct {
	Name      string `yaml:"name"`
	JwtSecret string `yaml:"jwt_secret"`
}

// ServerSection app.yaml server 配置
type ServerSection struct {
	HTTPPort int `yaml:"http_port"`
}

// DatabaseSection app.yaml database 配置
type DatabaseSection struct {
	Dialect    string `yaml:"dialect"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Dbname     string `yaml:"dbname"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Sslmode    string `yaml:"sslmode"`
	SqliteFile string `yaml:"sqlite_file"`
}

// RedisSection app.yaml redis 配置
type RedisSection struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// EmailSection app.yaml email 配置
type EmailSection struct {
	SmtpServer   string `yaml:"smtp_server"`
	SmtpPassword string `yaml:"smtp_password"`
	From         string `yaml:"from"`
}

// CodeSection app.yaml code 配置
type CodeSection struct {
	Min            int           `yaml:"min"`
	Max            int           `yaml:"max"`
	ValidityPeriod time.Duration `yaml:"validity_period"`
}

// 所有的配置项
var (
	App      AppSection
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
	Email    EmailSection
	Code     CodeSection
)

// app.yaml 对应结构体
var config struct {
	App      AppSection
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
	Email    EmailSection
	Code     CodeSection
}

func getCurrentPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func ReadValues(filenames ...string) (string, error) {
	if len(filenames) <= 0 {
		return "", errors.New("You must provide at least one filename for reading Values")
	}
	var resultValues map[string]interface{}
	for _, filename := range filenames {

		var override map[string]interface{}
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		if err := yaml.Unmarshal(bs, &override); err != nil {
			log.Println(err)
			continue
		}

		// check if is nil. This will only happen for the first filename
		if resultValues == nil {
			resultValues = override
		} else {
			for k, v := range override {
				resultValues[k] = v
			}
		}

	}
	bs, err := yaml.Marshal(resultValues)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(bs), nil
}

// init 初始化加载配置文件
func init() {
	var envPath string
	if os.Getenv(gin.EnvGinMode) != gin.TestMode {
		envPath = path.Join(getCurrentPath(), "../../config/app.yml")
	} else {
		envPath = path.Join(getCurrentPath(), "../../config/app.test.yml")
	}
	localEnvPath := path.Join(getCurrentPath(), "../../config/app.local1.yml")
	fmt.Printf("Load config file '%s'\n", envPath)
	// 解析 app.yml
	file, err := ReadValues(envPath, localEnvPath)
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal([]byte(file), &config); err != nil {
		panic(err)
	}

	App = config.App
	Server = config.Server
	Database = config.Database
	Redis = config.Redis
	Email = config.Email
	Code = config.Code
}
