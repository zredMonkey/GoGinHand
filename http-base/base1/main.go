package main

import (
	"fmt"
	"log"
	"net/http"
)

/**
设置了2个路由，/和/hello，分别绑定 indexHandler 和 helloHandler ，
根据不同的HTTP请求会调用不同的处理函数。
访问/，响应是URL.Path = /，而/hello的响应则是请求头(header)中的键值对信息
*/

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("hello", helloHandler)
	// 第一个参数监听9999端口，第二个参数代表处理所有的HTTP请求
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// 响应是URL.Path = /
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// 响应则是请求头(header)中的键值对信息
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
