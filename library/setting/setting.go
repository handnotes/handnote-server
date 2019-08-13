package setting

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// ServerSection app.yaml server 配置.
type ServerSection struct {
	HTTPPort int `yaml:"http_port"`
}

// DatabaseSection app.yaml database 配置.
type DatabaseSection struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Sslmode  string `yaml:"sslmode"`
}

// RedisSection app.yaml redis 配置.
type RedisSection struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// 所有的配置项.
var (
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
)

// app.yaml 对应结构体.
var config struct {
	Server   ServerSection
	Database DatabaseSection
	Redis    RedisSection
}

// init 初始化加载配置文件.
func init() {
	// 解析 app.yml
	file, err := ioutil.ReadFile("config/app.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal([]byte(file), &config)
	Server = config.Server
	Database = config.Database
	Redis = config.Redis
	if err != nil {
		log.Fatalln(err)
	}
}
