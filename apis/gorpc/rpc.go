package gorpc

import (
	infra "github.com/masterzcw/resk.wen"
	"github.com/masterzcw/resk.wen/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc)) // 注册一个RPC结构体实例到
}
