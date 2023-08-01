package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	// 创建Gee实例
	r := gee.New()

	// 添加路由(静态路由)
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动WEB服务
	r.Run(":9999")
}
