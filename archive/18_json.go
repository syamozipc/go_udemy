package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	// jsonをdescructする

	myJson := `
[
    {
        "first_name": "Clark",
        "last_name": "Kent",
        "hair_color": "black",
        "has_dog": true
    },
    {
        "first_name": "Bruce",
        "last_name": "Wayne",
        "hair_color": "black",
        "has_dog": false
    }
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)

	if err != nil {
		log.Println("error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	// structをjsonにする

	var mySlice []Person

	var m1 Person
	m1.FirstName = "ryota"
	m1.LastName = "saito"
	m1.HairColor = "red"
	m1.HasDog = true

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "saaya"
	m2.LastName = "morita"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	// indentをつける（第2引数はインデントの前につけるprefix、第3引数は任意のインデントスペース）
	// prod環境はMarshalで良い
	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshallling", err)
	}

	fmt.Println(string(newJson))
}
