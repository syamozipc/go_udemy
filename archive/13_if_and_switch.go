package main

import "log"

func main() {
	myVar := "doeg"

	switch myVar {
	case "cat":
		log.Print("cat")
	case "dog":
		log.Print("dog")
	case "fish":
		log.Print("fish")

	default:
		log.Print("wrong")
	}

	if true {

	} else if false {

	} else {

	}
}
