package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http requestを受け取った際に実行される処理
	// ブラウザで http://localhost:8080 にアクセスすると実行される
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w,"hello world!")
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		fmt.Printf("Number of bytes written: %d", n)
	})

	// serverを http://localhost:8080 に立ち上げる
	// 引数を_にすることで、errorがreturnされた場合、それをスルーする
	_ = http.ListenAndServe(":8080", nil)
}