package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Person person。
type Person struct {
	ID      int    `json:"id"` // プロパティの大文字と小文字にシンタックス的違いがある。マジ？小文字だと死ぬ。
	Name    string `json:"name"`
	Email   string `json:"email"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Memo    string `json:"memo"`
}

func main() {
	person := Person{
		ID:      8,
		Name:    "John",
		Email:   "mail@yahoo.co.jp",
		Age:     21,
		Address: "Tokyo, Shinjuku",
		Memo:    "This is memo.",
	}

	fmt.Println(person)

	personJSON, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(personJSON))

	data := new(Person)
	err = json.Unmarshal(personJSON, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(*data)

	file, err := os.Create("./file.json")
	if err != nil {
		log.Fatal(err)
	}

	// defer は関数が終了する際に実行される。
	defer file.Close()

	_, err = file.Write(personJSON)

	if err != nil {
		log.Fatal(err)
	}
}
