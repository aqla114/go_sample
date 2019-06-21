package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("./test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// slice, array, map は make で作れる。 cf) struct とかは new
	message := make([]byte, 20)

	res, err := file.Read(message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(message), res)
}
