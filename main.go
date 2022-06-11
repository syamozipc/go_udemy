package main

import "log"

func main(){
	var myString string
	var myInt int

	myString = "hi"
	myInt = 11

	mySecondString := "another string"

	log.Println(myString, mySecondString, myInt)
}