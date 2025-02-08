package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("name cannot be empty")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// now returns a map whose key is string and val is string
func Hellos(names []string) (map[string]string, error) {
	// map to associate names with messages
	messages := make(map[string]string)

	// loop through names passed in
	// call hello for each
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in map, assocaite retrieved messages with name
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hello, %s!",
		"Hi, %s.",
		"Hey, %s!",
	}
	return formats[rand.Intn(len(formats))]
}
