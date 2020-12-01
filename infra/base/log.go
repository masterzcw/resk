package base

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func init() {
	// 定义日志格式
	// formater := &log.TextFormatter{} // 原生日志
	formater := &prefixed.TextFormatter{}            // 第三方日志, 这个貌似漂亮一点
	formater.FullTimestamp = true                    // 开始日志时间显示
	formater.TimestampFormat = "2006-01-02 03:04:05" // 年月日时分秒都是按数字记录的, 这比较怪
	formater.ForceFormatting = true                  // 强制格式化, 没看出来有什么卵用
	formater.SetColorScheme(&prefixed.ColorScheme{
		InfoLevelStyle: "green",
		WarnLevelStyle: "yellow",
		// TimestampStyle: "37",
		//...
	}) // 设置颜色
	formater.ForceColors = true    // 控制台高亮显示
	formater.DisableColors = false // 控制台高亮显示
	log.SetFormatter(formater)

	// 日志级别
	level := os.Getenv("log.debug")
	if level == "true" || true {
		log.SetLevel(log.DebugLevel)
	}

	log.Info("测试Info")
	log.Debug("测试Debug")

	/*
		日志文件和滚动配置
		github.com/lestrrat/go-file-rotatelogs
	*/

}
