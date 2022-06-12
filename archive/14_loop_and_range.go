package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}
}

func main() {
	animals := []string{"dog", "fish", "horse"}

	// current iteration, value
	// iは無視できる（_で定義）
	for _, animal := range animals {
		// log.Println(i, animal)
		log.Println(animal)
	}
}

func main() {
	animals := make(map[string]string)
	animals["dog"] = "fido"
	animals["cat"] = "fluffy"

	// key, value
	for animalType, animal := range animals {
		// log.Println(i, animal)
		log.Println(animalType, animal)
	}
}

func main() {
	// 文字列に対しても可能だが、runeになる
	var firstLine = "Once upon a midnight dreary"

	// key, value
	for i, l := range firstLine {
		log.Println(i, ":", l)// 0:79, 1:110...
	}
}

func main() {
	type User struct {
		FirstName string
		LastName string
		Email string
		Age int
	}

	var users []User
	users = append(users, User{"ryota", "saito", "syamozipc@gmail.com", 30})
	users = append(users, User{"saaya", "morita", "saaya@gmail.com", 26})

	// key, value
	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}