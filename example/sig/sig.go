package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// 启动一个http服务
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello,world!"))
	})
	server := &http.Server{Addr: ":8080", Handler: mux}
	go func() {
		fmt.Println(server.ListenAndServe())
	}()

	// 接收监听系统传入的信号的通道
	sigs := make(chan os.Signal)
	// 监听信号量
	signal.Notify(sigs)
	// 从通道中接收信号, 无可用信号是阻塞等待
	c := <-sigs
	fmt.Println(c.String())
	fmt.Println(server.Close())
	time.Sleep(20 * time.Second)
	fmt.Println("退出.")
}
