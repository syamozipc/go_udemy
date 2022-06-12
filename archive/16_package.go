package main

import (
	"log"

	"github.com/syamozipc/my_nice_program/helpers"
)

func main(){
	var myVar helpers.SomeType
	myVar.TypeName = "ryota"

	log.Println(myVar.TypeName)
}
