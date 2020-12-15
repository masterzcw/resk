package main

import (
	"time"

	"github.com/tietang/go-eureka-client/eureka"
)

func main1() {
	cfg := eureka.Config{
		DialTimeout: time.Second * 10, // 超时时间
	}
	client := eureka.NewClientByConfig([]string{
		"http://127.0.0.1:8761/eureka",
	}, cfg)
	// 创建实例对象, 设置当前应用实例数据. 包括主机名称, 应用名称, 实例IP, 端口, 续约周期, 是否SSL
	appName := "Go-Example"
	instance := eureka.NewInstanceInfo("test.com", appName, "127.0.0.2", 8080, 30, false)
	client.RegisterInstance(appName, instance) // 注册服务
	client.Start()                             // 开始心跳
	c := make(chan int, 1)
	<-c
}
