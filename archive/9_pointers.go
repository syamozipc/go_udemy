package main

import "log"

func main(){
	var myString string
	myString = "Green"

	log.Println("my string is set to", myString)// 2022/06/10 10:12:16 my string is set to Green

	// ポインター渡し
	changeUsingPointer(&myString)

	log.Println("after func call", myString)// after func call Red
}

// ポインターの概念（PHPの参照渡しのようなもの）
func changeUsingPointer(s *string) {
	// log.Panicln("s is set to ", s)// 2022/06/10 10:12:16 s is set to  0xc000010260

	newValue := "Red"
	*s = newValue
}