package main

import (
	"fmt"
	"net/http"
)

func main() {
	// /hello にアクセスされたら Hello, World! を返すルーティングを定義する
	http.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!\r\n")
	})

	// 8080 番ポートを listen
	_ = http.ListenAndServe(":8080", nil)
}
