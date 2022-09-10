package controller

import (
	"net/http"
)

// 外部パッケージに公開するインタフェース
type Router interface {
	HandleTodosRequest(w http.ResponseWriter, r *http.Request)
}

// 非公開のRouter構造体
type router struct {
	tc TodoController
}

func NewRouter(tc TodoController) Router {
	return &router{tc}
}

func (ro *router) HandleTodosRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetTodos(w, r)
	case "POST":
		ro.tc.PostTodo(w, r)
	case "PUT":
		ro.tc.PutTodo(w, r)
	case "DELETE":
		ro.tc.DeleteTodo(w, r)
	default:
		w.WriteHeader(405)
	}
}
