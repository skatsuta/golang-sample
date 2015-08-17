package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// /hello にアクセスされたら Hello, World! を返すルーティングを定義する
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "{\"foo\":\"bar\", \"baz\":[1,2,3]}\r\n")
	})

	// 8080 番ポートを listen
	log.Fatal(http.ListenAndServe(":8080", nil))
}
