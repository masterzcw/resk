package main

import (
	"fmt"
	"time"

	"github.com/tietang/props/ini"
)

func main() {
	conf := ini.NewIniFileConfigSource("config.ini")              // 指定配置文件
	fmt.Println(conf.GetIntDefault("app.server.port", 18080))     // 读取整型
	fmt.Println(conf.GetDefault("app.server.port", "18080"))      // 读取字符串
	fmt.Println(conf.GetBoolDefault("app.enabled", true))         // 读取Bool
	fmt.Println(conf.GetDurationDefault("app.time", time.Second)) // 读取时间

}
