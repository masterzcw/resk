package gorpc

import (
	"resk.com/infra"
	"resk.com/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc)) // 注册一个RPC结构体实例到
}
