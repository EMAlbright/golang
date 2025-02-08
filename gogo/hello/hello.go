package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set properties of predefined logger
	// entry refix and flag to diable rinting
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// make array of names
	names := []string{"Ethan", "Bob", "Billy"}
	// request a greeting message
	message, err := greetings.Hellos(names)
	// if error, print it
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
