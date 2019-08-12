package library

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

// AppConfig 项目主配置结构.
type AppConfig struct {
	HTTPPort string `yaml:"http_port"`
}

// 所有的配置项.
var (
	App AppConfig
)

// init 初始化加载配置文件.
func init() {
	// 解析 app.yml.
	file, err := ioutil.ReadFile("config/app.yml")
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal([]byte(file), &App)
	if err != nil {
		log.Fatalln(err)
	}
}
