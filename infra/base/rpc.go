package base

import (
	"net"
	"net/rpc"
	"reflect"

	log "github.com/sirupsen/logrus"
	"resk.com/infra"
)

var rpcServer *rpc.Server

func RpcServer() *rpc.Server {
	Check(rpcServer)
	return rpcServer
}
func RpcRegister(ri interface{}) {

	// 日志
	typ := reflect.TypeOf(ri)
	log.Infof("goRPC Register: %s", typ.String())

	// 注册
	RpcServer().Register(ri)
}

type GoRPCStarter struct {
	infra.BaseStarter
	server *rpc.Server
}

func (s *GoRPCStarter) Init(ctx infra.StarterContext) {
	s.server = rpc.NewServer()
	rpcServer = s.server
}
func (s *GoRPCStarter) Start(ctx infra.StarterContext) {

	port := ctx.Props().GetDefault("app.rpc.port", "8082")
	//监听网络端口
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Panic(err)
	}
	log.Info("tcp port listened for rpc:", port)
	//处理网络连接和请求
	go s.server.Accept(listener)
}
