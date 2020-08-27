package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("web/static/"))

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "이것은 고랑 기반 서버이다. 당신의 요청 정보: %s", request.URL.Path)
	})

	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8081", nil)
}
