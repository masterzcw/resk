package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, r *http.Request) {
		s := r.URL.RawQuery // url中?后面的部分
		writer.Write([]byte("hello,worlsd." + s))
	}) // 创建指定路由的服务
	http.ListenAndServe(":8082", nil) // 开放服务到制定端口
}
