package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//type middleWareHandler struct {
//	r *httprouter.Router
//}
//
//
//func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
//	m := middleWareHandler{}
//	m.r = r
//	return m
//}
//
//func (m middleWareHandler) ServerHTTP(res http.ResponseWriter,req *http.Request) {
//	// add we need middleware
//
//	m.r.ServeHTTP(res,req)
//}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", addUser)
	//router.POST("user/:user_name")
	return router
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}
func main() {
	r := RegisterHandlers()
	//mR := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000",r)
}
