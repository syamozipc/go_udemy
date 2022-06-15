package main

import (
	"fmt"
	"log"
	"net/http"

	// 他のpackageを読み込む（go modのmodule名とpackage名に対応している）
	"github.com/syamozipc/go_udemy/config"
	"github.com/syamozipc/go_udemy/pkg/handlers"
	"github.com/syamozipc/go_udemy/pkg/render"
)

const portNumber = ":8080"

// main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	// http requestを受け取った際に実行される処理
	// ブラウザで http://localhost:8080 にアクセスすると実行される
	// TODO: 何故.Repoがつくのか分からない
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("starting application on port %s", portNumber)

	// serverを http://localhost:8080 に立ち上げる
	// 引数を_にすることで、errorがreturnされた場合、それをスルーする
	_ = http.ListenAndServe(portNumber, nil)
}
