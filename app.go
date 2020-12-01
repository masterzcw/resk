package resk

import (
	"resk/infra"
	"resk/infra/base"
)

func init() {
	// 注册基础设施
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.IrisServerStarter{}) // 这必须是最后一个, 因为在机制中, 给路由是朱携程阻塞的
}
