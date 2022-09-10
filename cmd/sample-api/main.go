package main

import (
	"log"
	"net/http"

	"github.com/koga456/sample-api/controller"
	"github.com/koga456/sample-api/model/repository"
)

var tr = repository.NewTodoRepository()
var tc = controller.NewTodoController(tr)
var ro = controller.NewRouter(tc)

func main() {
	log.Print("test aaa")
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/todos/", ro.HandleTodosRequest)
	server.ListenAndServe()
}
