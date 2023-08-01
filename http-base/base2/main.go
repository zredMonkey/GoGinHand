package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 定义 Engine 结构体  对所有request的handler
type Engine struct{}

/**
源码：
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
Handler是一个接口，需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了 ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了
*/

// 第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应
// 第二个参数包含了该HTTP请求的所有信息
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
