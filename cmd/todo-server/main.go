package main

import (
	"../../internal"
	"../../internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func main() {
	router := mux.NewRouter()

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

		fmt.Print("hello")

		todoItem := TodoService.GetTodo(id)
		json.NewEncoder(writer).Encode(todoItem)
	})

	router.HandleFunc("/newitem", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Print("hello new")
		decoder := json.NewDecoder(request.Body)
		var requestItem models.NewTodoItem

		err := decoder.Decode(&requestItem)

		if err != nil {
			panic(err)
		}

		fmt.Print(requestItem)
		TodoService.AddTodo(&requestItem)
	}).Methods("POST")

	router.
		PathPrefix("/static").
		Handler(http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))

	http.ListenAndServe(":8081", router)
}
