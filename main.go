// 1行目は必ずpackage declaration
package main

import "fmt"

// main()は引数を取らない
func main(){
	fmt.Println("hello, world")

	var whatToSay string
	// 64bitのPCを使用しているので
	var i int64

	whatToSay = "goodby cool world"

	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to ", i)

	// 型指定のショートカット
	whatWasSaid, theOtherThing := saySomething()

	fmt.Println("the function returned", whatWasSaid, theOtherThing)
}

func saySomething() (string, string) {
	return "something", "else"
}