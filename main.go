package main

import "log"

func main(){
	var myString string
	myString = "Green"

	log.Println("my string is set to", myString)

	// ポインター渡し
	changeUsingPointer(&myString)

	log.Println("after func call", myString)
}

// ポインターの概念（PHPの参照渡しのようなもの）
func changeUsingPointer(s *string) {
	log.Panicln("s is set to ", s)

	newValue := "Red"
	*s = newValue
}