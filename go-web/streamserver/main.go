package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
	l *ConnLimiter
}

func NewMiddleWareHandler(r *httprouter.Router, cc int) http.Handler {
	m := middleWareHandler{}
	m.r = r
	m.l = NewConnLimiter(cc)
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//如果没有拿到 token（不能获取连接）
	if !m.l.getConn() {
		sendErrorResponse(w, http.StatusTooManyRequests, "Too many requests")
		return
	}

	m.r.ServeHTTP(w, r)
	//施放token
	defer m.l.ReleaseConn()
}
func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id",streamHandler)
	router.POST("/upload/:vid-id",uploadHandler)
	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r, 2)
	http.ListenAndServe(":9000",mh)
}