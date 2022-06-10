// 1行目は必ずpackage declaration
package main

import (
	"log"
	"time"
)

// type名もproperth名も大文字始まりにしないと、main()内で呼び出せない
type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

// main()は引数を取らない
func main(){
	user := User {
		FirstName: "Ryota",
		LastName: "Saito",
		PhoneNumber: "03-1111-2222",
	}

	log.Println(
		user.FirstName,
		user.LastName,
		"birthdate:",
		user.BirthDate,
	)// Ryota Saito birthdate: 0001-01-01 00:00:00 +0000 UTC
}