package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"../../internal"
	"strconv"
)

func main() {
	router := mux.NewRouter()
	fs := http.FileServer(http.Dir("web/static/"))

	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "이것은 고랑 기반 서버이다. 당신의 요청 정보: %s", request.URL.Path)
	})

	router.HandleFunc("/hello/{name}", func(writer http.ResponseWriter, request *http.Request) {
		vars := mux.Vars(request)
		fmt.Fprintf(writer, "%s야, 반갑다. 즐거운 코딩 생활 되길 바란다.", vars["name"])
	})

	router.HandleFunc("/todo/{id}", func(writer http.ResponseWriter, request *http.Request) {
		idString := mux.Vars(request)["id"]
		id, _ := strconv.Atoi(idString)

		todoItem := TodoService.GetTodo(id)
		json.NewEncoder(writer).Encode(todoItem)
	})

	router.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8081", router)
}
