package base

import (
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	log "github.com/sirupsen/logrus"
	"resk.com/infra"
)

var callbacks []func()

func Register(fn func()) {
	callbacks = append(callbacks, fn)
}

type HookStarter struct {
	infra.BaseStarter
}

func (s *HookStarter) Init(ctx infra.StarterContext) {
	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM) // 注册信号量
	go func() {
		for {
			c := <-sigs // 监听信号量
			log.Info("notify: ", c)
			for _, fn := range callbacks {
				fn()
			}
			fmt.Println("退出.")
			break
			os.Exit(0)
		}
	}()

}

func (s *HookStarter) Start(ctx infra.StarterContext) {
	starters := infra.GetStarters() // 获得所有的注册项

	for _, s := range starters {
		typ := reflect.TypeOf(s)
		log.Infof("【Register Notify Stop】:%s.Stop()", typ.String())
		// 注册stop
		Register(func() {
			s.Stop(ctx)
		})
	}

}
