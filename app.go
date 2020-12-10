package resk

import (
	"fmt"

	"resk.com/apis/gorpc"
	_ "resk.com/apis/web"
	_ "resk.com/core/accounts"
	_ "resk.com/core/envelopes"
	"resk.com/infra"
	"resk.com/infra/base"
	"resk.com/jobs"
)

func init() {
	fmt.Println("/Users/crr/golang/resk.com/app.go->init()")
	// 注册基础设施
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.HookStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{}) // 这必须是最后一个, 因为在机制中, 给路由是朱携程阻塞的
}
