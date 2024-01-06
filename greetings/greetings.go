package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Please provide name in Hello() function call")
	}

	message := fmt.Sprintf(randomMessageFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if (err != nil) {
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

func randomMessageFormat() string {
	formats := []string{
		"Hi, %v Welcome!",
		"Great to see u %v",
		"Buenos Dias, %v",
	}

	return formats[rand.Intn(len(formats))]
}