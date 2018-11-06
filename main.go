package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)


type middleWareHandler struct { // 继承了路由
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

// middleWareHandler重写了serverHTTP
func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	validateUserSession(r) // authMiddleWare
	m.r.ServeHTTP(w, r)
	println("end")
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", addUser)

	return router
}

func main() {
	r := RegisterHandlers() //路由注册
	mh := NewMiddleWareHandler(r) // 重写ServerHttp
	http.ListenAndServe(":8000", mh)
}