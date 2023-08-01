package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc 类型，给框架用户用的，用来定义路由映射的处理方法
type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	// 路由映射表 key 由请求方法和静态路由地址构成，例如GET-/、GET-/hello、POST-/hello
	// 这样针对相同的路由，如果请求方法不同,可以映射不同的处理方法(Handler)
	// value 是用户映射的处理方法
	router map[string]HandlerFunc
}

// 构造方法
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// 会将路由和处理方法注册到映射表 router 中
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 是 ListenAndServe 的包装
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 作用：解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
