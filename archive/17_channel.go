package main

import (
	"log"

	"github.com/syamozipc/my_nice_program/helpers"
)

const numPool = 1000
func CalclateValue(intChan chan int){
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main(){
	// intのみを保持するchannel
	intChan := make(chan int)

	// defer横の処理は、今のfuncが実行終了後に実行される
	// main()実行完了後、channelが閉じられる
	defer close(intChan)

	// go routineを呼び出し、channelを実行する
	go CalclateValue(intChan)

	// channelをlistenする
	num := <- intChan

	log.Println(num)
}