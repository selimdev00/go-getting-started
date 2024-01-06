package main

import (
	"fmt"

	"rsc.io/quote/v4"

	"example.com/directed-message"
)

func main() {
	message := directedmessage.DirectedMessage("Selim", quote.Go())

	fmt.Println(message)
}