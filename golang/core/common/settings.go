package common

import (
	"github.com/go-ini/ini"
	"log"
	"os"
)

//ini配置文件
var cfg *ini.File

//服务器配置
var ServerSetting = &server{}

//初始化方法
func Init() {

	var err error

	dir, _ := os.Getwd()
	env := os.Getenv("GO_ENVIRONMENT")

	switch env {
	case "production":
		cfg, err = ini.Load(dir + "/conf/app.prod.ini")
		break
	default:
		cfg, err = ini.Load(dir + "/conf/app.dev.ini")
		break
	}

	if err != nil {
		panic(err)
	}

	mapTo("server", ServerSetting)

}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
