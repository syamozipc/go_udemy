package main

import (
	"log"
)

type User struct {
	FirstName string
	LastName string
}

func main(){
	// mapについて
	// - PHPやJSでは、mapでは順番通りに要素が格納されるが、Goではランダムに並び替えられる
    // - そのため、必ずkeyを指定し、取得時はそのkeyを指定する必要がある
	// [keyの型]valueの型
	myMap := make(map[string]User)

	me := User {
		FirstName: "ryota",
		LastName: "saito",
	}

	myMap["me"] = me

	log.Println(myMap) // map[me:{ryota saito}]
}


func main(){
	// sliceの宣言
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	// sortは破壊的メソッドっぽい
	sort.Ints(mySlice)

	log.Println(mySlice)// [1 2 3]
}


func main(){
	// sliceの宣言（ショートハンド）
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	log.Println(numbers)// [1 2 3 4 5 6 7 8 9 10]

	// 6番目から9番目の前まで
	log.Println(numbers[6:9])// [7 8 9]
}

func main(){
	names := []string{"one", "seven", "fish", "cat"}

	log.Println(names)// [one seven fish cat]
}