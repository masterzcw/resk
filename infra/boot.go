package infra

import (
	"reflect"

	"github.com/prometheus/common/log"
	"github.com/tietang/props/kvs"
)

// 应用程序启动管理器
type BootApplication struct {
	IsTest     bool
	conf       kvs.ConfigSource
	starterCtx StarterContext
}

// 创建注册机制
func New(conf kvs.ConfigSource) *BootApplication {
	e := &BootApplication{conf: conf, starterCtx: StarterContext{}}
	e.starterCtx.SetProps(conf)
	return e
}

func (b *BootApplication) Start() {
	//1. 初始化
	b.init()
	//2. 安装
	b.setup()
	//3. 启动
	b.start()
}

//程序初始化
func (e *BootApplication) init() {
	log.Info("Initializing starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debugf("Initializing: PriorityGroup=%d,Priority=%d,type=%s", v.PriorityGroup(), v.Priority(), typ.String())
		v.Init(e.starterCtx)
	}
}

//程序安装
func (e *BootApplication) setup() {
	log.Info("Setup starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debug("Setup: ", typ.String())
		v.Setup(e.starterCtx)
	}
}

//程序开始运行，开始接受调用
func (b *BootApplication) start() {
	log.Info("Starting starters...")
	for i, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debug("Starting: ", typ.String())
		if v.StartBlocking() {
			// go starter.Start(b.starterContext) // 阻塞的, 另开协程
			//如果是最后一个可阻塞的，直接启动并阻塞
			if i+1 == len(GetStarters()) {
				v.Start(b.starterCtx)
			} else {
				//如果不是，使用goroutine来异步启动，
				// 防止阻塞后面starter
				go v.Start(b.starterCtx)
			}
		} else {
			v.Start(b.starterCtx)
		}
	}
}

//程序开始运行，开始接受调用
func (e *BootApplication) Stop() {

	log.Info("Stoping starters...")
	for _, v := range GetStarters() {
		typ := reflect.TypeOf(v)
		log.Debug("Stoping: ", typ.String())
		v.Stop(e.starterCtx)
	}
}
