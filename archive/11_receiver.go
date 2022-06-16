package main

import "log"

type myStruct struct {
	FirstName string
}

// structのmethodとして機能する
func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main(){
	var myVar myStruct
	myVar.FirstName = "ryota"

	myVar2 := myStruct {
		FirstName: "saaya",
	}

	log.Println("myVar is set to", myVar.printFirstName())// myVar is set to ryota
	log.Println("myVar2 is set to", myVar2.printFirstName())// myVar2 is set to saaya
}