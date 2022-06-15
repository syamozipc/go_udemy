package main

import (
	"fmt"
	"net/http"

	// 他のpackageを読み込む（go modのmodule名とpackage名に対応している）
	"github.com/syamozipc/go_udemy/pkg/handlers"
)

const portNumber = ":8080"

// main application function
func main() {
	// http requestを受け取った際に実行される処理
	// ブラウザで http://localhost:8080 にアクセスすると実行される
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("starting application on port %s", portNumber)

	// serverを http://localhost:8080 に立ち上げる
	// 引数を_にすることで、errorがreturnされた場合、それをスルーする
	_ = http.ListenAndServe(portNumber, nil)
}
