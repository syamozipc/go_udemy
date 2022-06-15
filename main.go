package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addvalues(2, 2)
	_, _ = fmt.Fprintf(w, "this is the about page and 2 + 2 is %d", sum)
}

// addValues adds two integers and return the sum
func addvalues(x, y int) int {
	return x + y
}

// main application function
func main() {
	// http requestを受け取った際に実行される処理
	// ブラウザで http://localhost:8080 にアクセスすると実行される
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("starting application on port %s", portNumber)

	// serverを http://localhost:8080 に立ち上げる
	// 引数を_にすることで、errorがreturnされた場合、それをスルーする
	_ = http.ListenAndServe(portNumber, nil)
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := DivideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")

		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", 100.0, 0.0, f)
}

func DivideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot devide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}
