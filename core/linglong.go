package core
// 在时间的长河中砥砺前行

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	routeKeySep = "-"
)

// HandlerFunc
// @description 定义了请求的处理方法
// @param w http.ResponseWriter 响应体接口对象，实现了http.ResponseWriter接口的实例
// @param r *http.Request 请求体的结构体指针，结构体中包含了一个请求的method, url, protocol(HTTP/HTTPS), header等信息
type HandlerFunc func(w http.ResponseWriter, r *http.Request)

// Hope
// @description 框架的启动入口，实现Handler接口
type Hope struct {
	router map[string]HandlerFunc
}

// New
// @description return the constructor of *core.Hope
func New() *Hope {
	return &Hope{
		router: make(map[string]HandlerFunc),
	}
}

func concatRouterKey(method, pattern string) string {
	var builder strings.Builder
	builder.Grow(len(method) + len(pattern) + len(routeKeySep))
	builder.WriteString(method)
	builder.WriteString(routeKeySep)
	builder.WriteString(pattern)

	return builder.String()
}

// ServeHTTP
// @description 实现Handler接口中的ServeHTTP方法
func (hope *Hope) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := concatRouterKey(r.Method, r.URL.Path)
	if handler, ok := hope.router[key]; ok {
		fmt.Printf("new-request: %s %s\n", r.Method, r.URL.Path)
		handler(w, r)
	} else {
		// r.URL: 返回请求的有效转义形式
		fmt.Fprintf(w, "404 url not found: %s\n", r.URL)
	}
}

func (hope *Hope) Run(addr string) (err error) {
	fmt.Printf("%d registered requests:\n", len(hope.router))
	for reqInfo := range hope.router {
		info := strings.SplitN(reqInfo, routeKeySep, 2)
		fmt.Printf("[%s] %s\n", info[0], info[1])
	}
	return http.ListenAndServe(addr, hope)
}

// addRoute
// @description 注册请求路由及其handler
// @param method string 请求方法，如：get, post
// @param pattern string 请求的非转义url
// @param handler HandlerFunc 请求的业务逻辑处理方法
func (hope *Hope) addRoute(method, pattern string, handler HandlerFunc) {
	hope.router[concatRouterKey(method, pattern)] = handler
}

// GET
// @description 注册get请求
func (hope *Hope) GET(pattern string, handler HandlerFunc) {
	hope.addRoute("GET", pattern, handler)
}

// POST
// @description 注册post请求
func (hope *Hope) POST(pattern string, handler HandlerFunc) {
	hope.addRoute("POST", pattern, handler)
}
