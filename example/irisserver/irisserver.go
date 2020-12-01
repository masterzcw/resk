package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
)

func main() {
	app := iris.Default()
	app.Get("/hello", func(ctx iris.Context) {
		ctx.WriteString("hello,,,world!")
	})

	v1 := app.Party("/v1") // 定义路由分组
	v1.Use(func(context iris.Context) {
		logrus.Info("自定义中间件")
		context.Next()
	}) // 定义路由中间件

	v1.Get("/users/{id:uint64 min(2)}", func(ctx iris.Context) {
		id := ctx.Params().GetUint64Default("id", 0)
		ctx.WriteString(strconv.Itoa(int(id)))
	}) // 组内路由 /v1/users/123
	app.Get("/users/{action:path}", func(ctx iris.Context) {
		a := ctx.Params().Get("action")
		ctx.WriteString("/users/{action:path}:" + a)
	})

	/*
		定义错误
		后定义的会覆盖先定义的, 所以后定义的通常粒度会比较细
	*/
	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.WriteString("看起来出错了!")
	})
	app.OnErrorCode(http.StatusNotFound, func(ctx iris.Context) {
		// app.OnErrorCode(http.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("访问路径不存在")
	})

	fmt.Println(app.Run(iris.Addr(":8081")))
	//http://localhost:8081/hello
}
