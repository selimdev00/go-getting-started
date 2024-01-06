package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)


func main() {
	log.SetPrefix("greetings: ")

	log.SetFlags(0)

	names := []string{"Selim", "Oraz", "Ata"}

	message, err := greetings.Hello("SELIM")
	if err != nil {
		log.Fatal(err)
	}

	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	fmt.Println(message)
}